// Copyright © 2017 EasyStack Inc. Shawn Wang <shawn.wang@easystack.cn>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package ostree

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"gitlab.com/EasyStack/yakety/lib/env"
	"gitlab.com/EasyStack/yakety/lib/errors"
	"gitlab.com/EasyStack/yakety/lib/utils"
)

type BackupDeployment struct {
	Version   string
	Origin    string
	Osname    string
	Commit    string `json:"base-checksum"`
	Checksum  string
	Id        string
	Timestamp uint64 `json:"timestamp"`
}

type backupData struct {
	Repo    string             `yaml:"repo"`
	Backups []BackupDeployment `yaml:"backups,flow"`
}

type rpmOstreeDeployment struct {
	BackupDeployment
	Booted bool
}

type rpmOstreeStatusOutput struct {
	Deployments []rpmOstreeDeployment
	Transaction interface{}
}

var execCommand = exec.Command

func IsOstreeHost() bool {
	_, err := execCommand("rpm-ostree", "status", "--json").Output()
	if err == nil {
		return true
	} else {
		log.Println(errors.HOST_NOT_SUPPORT)
		return false
	}
}

func BackupPath() string {
	path := filepath.Join(env.YakRoot(), env.DataDir, "atomic")
	os.MkdirAll(path, 0755)
	return path
}

func BackupIndexFile() string {
	file := filepath.Join(BackupPath(), "backup.yml")
	return file
}

func configDiff() []string {
	var result []string
	out, err := exec.Command("ostree", "admin", "config-diff").Output()
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, each := range data {
		if string(each[0]) != "D" {
			file := each[5:len(each)]
			result = append(result, "/etc/"+file)
		}
	}
	return result
}

var sourcePath string
var targetPath string

func backupFileFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Print(err)
		return nil
	}

	base := strings.Split(sourcePath, `/`)
	ppath := strings.Split(path, sourcePath)

	tPath := filepath.Join(targetPath, base[len(base)-1], ppath[1])

	if info.IsDir() {
		os.MkdirAll(tPath, info.Mode())
	} else {
		err := os.Link(path, tPath)
		if err != nil {
			log.Printf("ln %s %s :%s \n", path, tPath, err)
		}
	}
	return nil
}

func Backup(repoName string) error {
	status := getRpmOstreeStatus()
	deployment := getCurrentDeployment(status)

	basePath := filepath.Join("/sysroot/ostree/deploy/", deployment.Osname, "deploy")
	// FIXME shouldn't hardcode env.YakRoot()
	rootPath := filepath.Join("/sysroot/ostree/deploy/", deployment.Osname, "var/roothome/yak", env.DataDir, "atomic")
	os.MkdirAll(rootPath, 0755)
	sourcePath = filepath.Join(basePath, strings.Split(deployment.Id, deployment.Osname+"-")[1], `etc`)
	targetPath = filepath.Join(rootPath, deployment.Checksum[0:6])

	updateBackupList(deployment.BackupDeployment, repoName)
	return filepath.Walk(sourcePath, backupFileFunc)
}

// Keep SaveDiffTarGz for reference,
// we might remove it later
func SaveDiffTarGz(repoName string) bool {
	status := getRpmOstreeStatus()
	deployment := getCurrentDeployment(status)
	name := deployment.Checksum[0:6]
	path := filepath.Join(BackupPath(), name+".tar.gz")
	files := configDiff()
	file, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	// set up the gzip writer
	gw := gzip.NewWriter(file)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	// grab the paths that need to be added in

	for _, i := range files {
		if err := utils.AddTarFile(tw, i); err != nil {
			log.Fatalln(err)
			return false
		}
	}
	updateBackupList(deployment.BackupDeployment, repoName)
	return true
}

func getRpmOstreeStatus() rpmOstreeStatusOutput {
	var status rpmOstreeStatusOutput
	out, _ := execCommand("rpm-ostree", "status", "--json").Output()
	json.Unmarshal(out, &status)
	return status
}

func getCurrentDeployment(status rpmOstreeStatusOutput) rpmOstreeDeployment {
	for _, d := range status.Deployments {
		log.Printf("%v %q \n\n", d, d)
		if d.Booted {
			return d
		}
	}
	return rpmOstreeDeployment{}
}

func GetBackupList() backupData {
	result := backupData{}
	file := BackupIndexFile()
	data, err := ioutil.ReadFile(file)
	if err == nil {
		yaml.Unmarshal([]byte(data), &result)
	}
	return result
}

func updateBackupList(b BackupDeployment, repoName string) backupData {
	file := BackupIndexFile()
	result := GetBackupList()
	new_item := true
	for _, v := range result.Backups {
		if v.Checksum == b.Checksum {
			new_item = false
		}
	}
	if new_item {
		result.Backups = append(result.Backups, b)
		result.Repo = repoName
		data, _ := yaml.Marshal(&result)
		ioutil.WriteFile(file, data, 0644)
	}
	return result
}

func GetRollbackDeployment(b string) BackupDeployment {
	items := GetBackupList()
	for i, v := range items.Backups {
		if v.Checksum[0:len(b)] == b {
			return v
		}
		if b == "" && i == (len(items.Backups)-1) {
			return v
		}
		fmt.Println(v.Checksum)
	}
	return BackupDeployment{}
}

func (b *BackupDeployment) Rollback(dryrun bool) bool {
	fmt.Printf("do backup %v\n", b)
	deployCmd := utils.Cmd("rpm-ostree", "deploy", b.Commit)
	utils.DoRun(deployCmd, dryrun)

	status := getRpmOstreeStatus()
	for _, v := range status.Deployments {
		if v.Commit == b.Commit {
			idPath := strings.Split(v.Id, v.Osname+"-")[1]
			targetPath = filepath.Join("/sysroot/ostree/deploy/", b.Osname, "deploy", idPath)
		}
	}

	files, err := ioutil.ReadDir(targetPath + "/etc/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		if !dryrun {
			fmt.Println("remove", targetPath+"/etc/"+file.Name())
			err = os.RemoveAll(targetPath + "/etc/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	sourcePath = filepath.Join("/sysroot/ostree/deploy/", b.Osname, "var/roothome/yak", env.DataDir, "atomic", b.Checksum[0:6], "etc")

	if !dryrun {
		filepath.Walk(sourcePath, backupFileFunc)
	}
	fmt.Printf("Rollbacked, please reboot. (systemctl reboot)\n")
	return true
}
