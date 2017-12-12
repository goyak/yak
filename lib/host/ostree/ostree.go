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

type backupDeployment struct {
	Version   string
	Origin    string
	Osname    string
	Commit    string `json:"base-checksum"`
	Checksum  string
	Timestamp uint64 `json:"timestamp"`
}

type backupData struct {
	Repo    string             `yaml:"repo"`
	Backups []backupDeployment `yaml:"backups,flow"`
}

type rpmOstreeDeployment struct {
	backupDeployment
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
	out, err := exec.Command("ostree", "admin", "config-diff").Output()
	if err != nil {

		log.Fatal(err)

	}
	data := fmt.Sprintf("%s", out)
	result := strings.Split(strings.TrimSpace(data), "\n")

	for index, each := range result {
		file := each[5:len(each)]
		result[index] = "/etc/" + file
	}
	return result
}

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
	updateBackupList(deployment.backupDeployment, repoName)
	return true
}

func getRpmOstreeStatus() rpmOstreeStatusOutput {
	var status rpmOstreeStatusOutput
	out, _ := execCommand("rpm-ostree", "status", "--json").Output()
	json.Unmarshal(out, &status)
	return status
}

func getCurrentDeployment(status rpmOstreeStatusOutput) rpmOstreeDeployment {
	fmt.Printf("getCurrentDeployment: %s\n", status)
	for _, d := range status.Deployments {
		fmt.Printf("%v %q \n\n", d, d)
		if d.Booted {
			return d
		}
	}
	return rpmOstreeDeployment{}
}

func getBackupList() backupData {
	var result backupData
	file := BackupIndexFile()
	data, _ := ioutil.ReadFile(file)
	yaml.Unmarshal([]byte(data), &result)
	return result
}

func updateBackupList(b backupDeployment, repoName string) backupData {
	file := BackupIndexFile()
	result := getBackupList()
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

func GetRollbackDeployment(b string) {
	items := getBackupList()
	// if b == "" {
	// 	return items[0].BackupDeployment
	// }
	for _, v := range items.Backups {
		fmt.Println(v.Checksum)
	}
}
