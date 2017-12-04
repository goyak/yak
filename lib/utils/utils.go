package utils

import (
	"io/ioutil"
	"log"

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
