package index

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type RecipeItem struct {
	Repo      string            // github.com/shawn111/abc
	Recipe    string            // abc/xxx
	Ref       string            `yaml:",omitempty"`
	Version   string            `yaml:",omitempty"`
	Installed bool              `yaml:",omitempty"`
	Extra     map[string]string `yaml:",inline"`
}

type Index struct {
	Remotes []string     `yaml:",flow"`
	Apps    []RecipeItem `yaml:",flow"`
}

func LoadIndex(file string) Index {
	index := Index{}
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
