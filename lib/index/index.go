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
