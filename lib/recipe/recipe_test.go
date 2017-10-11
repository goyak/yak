package recipe

import (
	"fmt"
	"log"
	"testing"

	"gopkg.in/yaml.v2"
)

var data = `
name: core
summary: Easystack Container Linux
version: '1.0'

# ${url}:${branch}@${from}
backend: ostree
source: http://mirror.centos.org/centos/7/atomic/x86_64/repo
branch: centos-atomic-host/7/x86_64/standard
backend: ostree
hash: 173278f2ccba80c5cdda4b9530e6f0388177fb6d27083dec9d61bbe40e22e064

description: |
  Easystack Container Linux

config:
  abc: zz
  abc2: zz
  abc1: zz
  abc3: zz
`

func TestRecipeConfig(t *testing.T) {
	r := RecipeConfig{}
	err := yaml.Unmarshal([]byte(data), &r)

	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)
}

func TestRecipeConfigDump(t *testing.T) {
	recipe := LoadRecipeConfig("atomic_yaml_testdata")
	recipe.Dump()
}
