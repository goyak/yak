package recipe

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gitlab.com/EasyStack/yakety/lib/env"
	"gitlab.com/EasyStack/yakety/lib/host/ostree"
)

type AtomicRecipeConfig struct {
	BaseRecipeConfig
}

func (r AtomicRecipeConfig) IsInstallable() bool {
	return ostree.IsOstreeHost()
}

func (r *AtomicRecipeConfig) BackupPath() string {
	path := filepath.Join(env.YakRoot(), env.DataDir, "atomic")
	os.MkdirAll(path, 0755)
	return path
}

func (r *AtomicRecipeConfig) createTarGz(name string, files []string) bool {
	// set up the output file
	// we might share atomic data cross different yak
	// $YAKPATH/index/atomic
	// normal recipe data dir
	// $YAKPATH/recipes/<r.Repo>/data
	// $YAKPATH/recipes/<r.Repo>/plugin
	path := filepath.Join(r.BackupPath(), name+".tar.gz")
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
		if err := addFile(tw, i); err != nil {
			log.Fatalln(err)
			return false
		}
	}
	return true
}

func addFile(tw *tar.Writer, path string) error {
	fmt.Printf("prepare: %s.\n", path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if stat, err := file.Stat(); err == nil {
		// now lets create the header as needed for this file within the tarball
		header := new(tar.Header)
		header.Name = path
		if !stat.IsDir() {
			header.Size = stat.Size()
			header.ModTime = stat.ModTime()
			header.Mode = int64(stat.Mode())
		}

		if err := tw.WriteHeader(header); err != nil {
			fmt.Printf(" >> n: %s \n", header.Name)
			fmt.Printf(" >> size: %d \n", header.Size)
			fmt.Printf(" >> m: %d \n", header.Mode)
			fmt.Printf(" >> err: %s \n", err)
			return err
		}

		if !stat.IsDir() {
			// write the header to the tarball archive
			// copy the file data to the tarball
			fmt.Printf("copying %s ... m: %d\n", path, header.Mode)
			if _, err := io.Copy(tw, file); err != nil {
				return err
			}
			fmt.Printf("copied %s.\n", path)
		}
	}
	return nil
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

func (r AtomicRecipeConfig) Install() bool {
	// backup current local config
	// ostree admin config-diff
	status := getRpmOstreeStatus()
	deployment := getCurrentDeployment(status)
	tarFileName := deployment.Checksum[0:6]
	files := configDiff()
	result := r.createTarGz(tarFileName, files)
	if !result {
		return false
	}

	file := filepath.Join(env.YakRoot(), env.IndexDir, "atomic", "backup.yml")
	deployment.backupDeployment.updateBackup(file)

	remoteName := strings.Split(r.Branch, "/")[0]
	addRemoteCmd := exec.Command("ostree", "remote", "add", "--if-not-exists", "--no-gpg-verify", remoteName, r.Source)
	fmt.Printf("::: %s  %q \n", addRemoteCmd.Path, addRemoteCmd.Args)
	// addRemoteCmd.Run()

	pullCmd := exec.Command("ostree", "pull", remoteName, r.Commit)
	fmt.Printf("::: %s  %q \n", pullCmd.Path, pullCmd.Args)
	// pullCmd.Run()

	// ostree admin deploy d518b37c348eb814093249f035ae852e7723840521b4bcb4a271a80b5988c44a
	// rpm-ostree deploy 173278f2ccba80c5cdda4b9530e6f0388177fb6d27083

	cmd := exec.Command("rpm-ostree", "deploy", r.Commit)
	cmd.Run()
	fmt.Printf("::: %s  %q \n", cmd.Path, cmd.Args)

	// Prepare to reboot
	// cmd := exec.Command("ostree", "/tmp/aaa")
	// cmd.Run()
	return true
}
