package index

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/EasyStack/yakety/lib/recipe"
	"gopkg.in/yaml.v2"
)

var data = `
remotes:
  - abc
  - xxx
apps:
  -
    name: github.com/foo/abc
    repo: abc/bar
    ref: asdasdas
    installed: true
  -
    name: github.com/foo2/abc2
    repo: abc2/barx
    ref: asdasdas
`

func TestIndex(t *testing.T) {
	idx := Index{}
	err := yaml.Unmarshal([]byte(data), &idx)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	assert.IsType(t, new(Index), &idx)
}

func TestLoadIdex(t *testing.T) {
	idx := LoadIndex("testdata/index.yaml")

	assert.IsType(t, new(Index), &idx)
	assert.Equal(t, idx.Name, "testdata/index.yaml")
}

func TestInstallApp(t *testing.T) {
	const file = `/tmp/test_index.yaml`
	idx := Index{
		Name: file,
	}
	app := recipe.RecipeConfig{
		Name: "abc",
	}

	assert.Equal(t, len(idx.Apps), 0)
	idx.Install(app)
	assert.Equal(t, len(idx.Apps), 1)
	os.Remove(file)
}
