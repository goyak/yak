package index

import (
	"gitlab.com/EasyStack/yakety/lib/recipe"
	"gitlab.com/EasyStack/yakety/lib/utils"
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
	var index Index
	utils.LoadYaml(file, &index)
	index.Name = file
	return index
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
	utils.SaveYaml(index.Name, &index)
}
