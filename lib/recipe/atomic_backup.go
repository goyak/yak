package recipe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"os/exec"
)

type backupDeployment struct {
	Version   string
	Origin    string
	Osname    string
	Commit    string `json:"base-checksum"`
	Checksum  string
	Timestamp uint64 `json:"timestamp"`
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

func (r *AtomicRecipeConfig) getBackupList() []backupDeployment {
	var result []backupDeployment
	file := r.backupIndexFile()
	data, _ := ioutil.ReadFile(file)
	yaml.Unmarshal([]byte(data), &result)
	return result
}

func (r *AtomicRecipeConfig) updateBackupList(b backupDeployment) []backupDeployment {
	file := r.backupIndexFile()
	result := r.getBackupList()
	new_item := true
	for _, v := range result {
		if v.Checksum == b.Checksum {
			new_item = false
		}
	}
	if new_item {
		result = append(result, b)

		data, _ := yaml.Marshal(&result)
		ioutil.WriteFile(file, data, 0644)
	}
	return result
}

func (b backupDeployment) updateBackup(file string) []backupDeployment {

	var result []backupDeployment

	data, _ := ioutil.ReadFile(file)
	yaml.Unmarshal([]byte(data), &result)
	new_item := true
	for _, v := range result {
		if v.Checksum == b.Checksum {
			new_item = false
		}
	}
	if new_item {
		result = append(result, b)

		data, _ = yaml.Marshal(&result)
		ioutil.WriteFile(file, data, 0644)
	}
	return result
}
