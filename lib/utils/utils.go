package utils

import (
	"io/ioutil"
	"log"
	"os/exec"

	"gopkg.in/yaml.v2"
)

func SaveYaml(file string, in interface{}) {
	d, err := yaml.Marshal(&in)
	if err != nil {
		log.Fatalf("error: %v", err)
		panic(err)
	}
	err = ioutil.WriteFile(file, d, 0644)
	if err != nil {
		log.Fatalf("save file: %s error: %v", file, err)
		panic(err)
	}
}

var Cmd = exec.Command

func DoRun(cmd *exec.Cmd, dryrun bool) error {
	log.Printf("::: %s  %q \n", cmd.Path, cmd.Args)
	if !dryrun {
		return cmd.Run()
	}
	return nil
}
