package index

import (
	"fmt"
	"io/ioutil"

	"github.com/goyak/yak/lib/recipe"
	"github.com/goyak/yak/lib/utils"
	"gopkg.in/yaml.v2"
)

type RecipeItem struct {
	recipe.RecipeConfig `yaml:",inline"`
	Installed           bool `yaml:",omitempty"`
}

type Index struct {
	Name    string
	Remotes []string
	Apps    []RecipeItem
}

func LoadIndex(file string) Index {
	var out Index
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("load file: %s\n", file)
		panic(err)
	}
	err = yaml.Unmarshal([]byte(data), &out)
	if err != nil {
		fmt.Printf("load file: %s\n", file)
		panic(err)
	}
	out.Name = file
	return out
}

func (index *Index) Install(app recipe.RecipeConfig) {
	installed := false
	appx := RecipeItem{
		RecipeConfig: app,
		Installed:    true,
	}

	for idx, _ := range index.Apps {
		if index.Apps[idx].Repo == app.Repo {
			index.Apps[idx] = appx
			installed = true
		}
	}
	if !installed {
		index.Apps = append(index.Apps, appx)
	}
	utils.SaveYaml(index.Name, &index)
}
