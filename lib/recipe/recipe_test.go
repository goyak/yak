package recipe

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
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
hash: 173278f2ccba80c5cdda4b9530e6f0388177fb6d27083dec9d61bbe40e22e064

description: |
  Easystack Container Linux

# extra configs
xx: cc
abc: zz
`

func TestRecipeConfig(t *testing.T) {
	rc := RecipeConfig{}
	err := yaml.Unmarshal([]byte(data), &rc)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	assert.IsType(t, new(RecipeConfig), &rc)
	assert.NotNil(t, rc.Name)
	assert.Equal(t, rc.Name, "core")
	assert.Equal(t, rc.Extra[`abc`], "zz")
}

func TestLoadRecipeConfig(t *testing.T) {
	r := LoadRecipeConfig("testdata/atomic.yml")

	assert.IsType(t, new(IRecipeConfig), &r)
}

func TestRecipeGetRecipeConfig(t *testing.T) {
	recipe := LoadRecipeConfig("testdata/atomic.yml")

	cfg := recipe.GetRecipeConfig()
	assert.Equal(t, cfg.Version, "1.0")
}

func TestRecipeConfigGetExtra(t *testing.T) {
	recipe := LoadRecipeConfig("testdata/atomic.yml")

	cfg := recipe.GetRecipeConfig()
	val := cfg.GetExtra(`new_item`, `test`)
	assert.Equal(t, val, "test")
}

func TestRecipeConfigGetExtraExists(t *testing.T) {
	recipe := LoadRecipeConfig("testdata/atomic.yml")

	cfg := recipe.GetRecipeConfig()
	val := cfg.GetExtra(`abc`, `test`)
	assert.Equal(t, val, "zz")
}
