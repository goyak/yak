package index

import (
	"fmt"
	"io/ioutil"
	"log"

	"gitlab.com/EasyStack/yakety/lib/recipe"
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
	index := Index{}
	index.Name = file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(data), &index)
	if err != nil {
		panic(err)
	}
	return index
}

func (index *Index) Save() {
	d, err := yaml.Marshal(&index)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = ioutil.WriteFile(index.Name, d, 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func (index *Index) Install(app recipe.RecipeConfig) {
	appx := RecipeItem{
		RecipeConfig: app,
		Installed:    true,
	}
	apps := index.Apps
	for i, a := range apps {
		if a.Repo == app.Repo {
			if len(index.Apps) > 1 {
				fmt.Printf("%d %q %v", i, a, a)
				if (len(index.Apps) - 1) == i {
					index.Apps = index.Apps[:i]
				} else {
					index.Apps = append(index.Apps[:i], index.Apps[i+1])
				}
			} else {
				index.Apps = []RecipeItem{}
			}
		}
	}
	index.Apps = append(index.Apps, appx)
	fmt.Printf("%q %v", index.Apps, index.Apps)
	index.Save()
}
