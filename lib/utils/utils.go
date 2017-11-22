package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadYaml(file string, out interface{}) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("load file: %s\n", file)
		panic(err)
	}
	err = yaml.Unmarshal([]byte(data), &out)
	if err != nil {
		panic(err)
	}
}

func SaveYaml(file string, in interface{}) {
	d, err := yaml.Marshal(&in)
	if err != nil {
		//log.Fatalf("error: %v", err)
		panic(err)
	}
	err = ioutil.WriteFile(file, d, 0644)
	if err != nil {
		fmt.Printf("save file: %s\n", file)
		//log.Fatalf("error: %v", err)
		panic(err)
	}
}
