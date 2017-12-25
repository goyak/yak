package index

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/goyak/yak/lib/recipe"
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

func TestLoadIndex(t *testing.T) {
	idx := LoadIndex("testdata/index.yaml")

	assert.IsType(t, new(Index), &idx)
	assert.Equal(t, "testdata/index.yaml", idx.Name)
	assert.Equal(t, 2, len(idx.Apps))
}

func TestInstallApp(t *testing.T) {
	const file = `/tmp/test_index.yaml`
	idx := Index{
		Name: file,
	}
	app := recipe.RecipeConfig{
		Name: "abc",
	}

	assert.Equal(t, 0, len(idx.Apps))
	idx.Install(app)
	assert.Equal(t, 1, len(idx.Apps))
	os.Remove(file)
}

func TestUpgradeApp(t *testing.T) {
	const file = `/tmp/test_index.yaml`
	idx := Index{
		Name: file,
	}
	idx.Install(recipe.RecipeConfig{
		Name:    "abc",
		Version: "1.0",
		Repo:    "ss/abc",
	})
	assert.Equal(t, 1, len(idx.Apps))
	assert.Equal(t, "1.0", idx.Apps[0].Version)

	idx.Install(recipe.RecipeConfig{
		Name:    "abc",
		Version: "2.0",
		Repo:    "ss/abc",
	})
	assert.Equal(t, 1, len(idx.Apps))
	assert.Equal(t, "2.0", idx.Apps[0].Version)

	idx.Install(recipe.RecipeConfig{
		Name:    "foo",
		Version: "2.0",
		Repo:    "ss/foo",
	})
	idx.Install(recipe.RecipeConfig{
		Name:    "bar",
		Version: "1.0",
		Repo:    "ss/bar",
	})
	assert.Equal(t, 3, len(idx.Apps))

	idx.Install(recipe.RecipeConfig{
		Name:    "abc",
		Version: "2.1",
		Repo:    "ss/abc",
	})

	assert.Equal(t, 3, len(idx.Apps))

	idx.Install(recipe.RecipeConfig{
		Name:    "foo",
		Version: "2.1",
		Repo:    "ss/foo",
	})
	assert.Equal(t, 3, len(idx.Apps))
	idx.Install(recipe.RecipeConfig{
		Name:    "bar",
		Version: "2.1",
		Repo:    "ss/bar",
	})
	assert.Equal(t, 3, len(idx.Apps))

	os.Remove(file)
}
