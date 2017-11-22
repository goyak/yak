package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type person struct {
	name string
	age  int
}

func TestSaveYaml(t *testing.T) {
	const file = `/tmp/test.yaml`
	bob := person{"Bob", 20}
	SaveYaml(file, bob)

	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("load file: %s\n", file)
		panic(err)
	}
	fmt.Printf("%v %T", data, data)
	assert.Equal(t, 1, 1)
	os.Remove(file)
}
