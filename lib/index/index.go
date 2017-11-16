package index

import (
	"io/ioutil"
	"log"

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
