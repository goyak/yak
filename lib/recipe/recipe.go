// Copyright © 2017 EasyStack Inc. Shawn Wang <shawn.wang@easystack.cn>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package recipe

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

type RecipeConfig struct {
	Name        string `yaml:"name"`
	Repo        string
	Summary     string
	Version     string
	Backend     string
	Source      string
	Branch      string
	Commit      string // Ref
	Description string
	Extra       map[string]string `yaml:",inline"`
}

func (cfg *RecipeConfig) GetExtra(key string, defaultValue string) string {
	if _, ok := cfg.Extra[key]; !ok {
		cfg.Extra[key] = defaultValue
	}
	return cfg.Extra[key]
}

type BaseRecipeConfig struct {
	RecipeConfig
	Path string
}

type IRecipeConfig interface {
	IsRecipe() bool
	GetRecipeConfig() RecipeConfig
	IsInstallable() bool
	Install() bool
	Dump() string
	Fetch(root string) bool
}

func LoadRecipeConfig(file string) IRecipeConfig {
	config := RecipeConfig{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		panic(err)
	}
	if config.Backend == "atomic" {
		r := AtomicRecipeConfig{}
		r.init(config)
		r.Path = file
		return r
	} else {
		r := BaseRecipeConfig{}
		r.init(config)
		r.Path = file
		return r
	}
}

func (r *BaseRecipeConfig) init(cfg RecipeConfig) {
	r.RecipeConfig = cfg
}

func (r BaseRecipeConfig) IsRecipe() bool {
	if r.Path == "" {
		return false
	} else {
		return true
	}
}

func (r BaseRecipeConfig) Install() bool {
	return false
}

func (r BaseRecipeConfig) Fetch(root string) bool {
	cmd := exec.Command("git", "clone", "https://"+r.Name+".git", root+"/recipes/"+r.Name)
	fmt.Printf("git clone https://%s\n", r.Name)
	err := cmd.Run()
	if err != nil {
		log.Printf("Failed.\ncmd: %s,\nerror: %v\n", strings.Join(cmd.Args, " "), err)
		return false
	}
	log.Printf("Done.\n")
	return true
}

func (r BaseRecipeConfig) GetRecipeConfig() RecipeConfig {
	return r.RecipeConfig
}

func (r BaseRecipeConfig) IsInstallable() bool {
	return true
}

func (r BaseRecipeConfig) Dump() string {
	d, err := yaml.Marshal(&r.RecipeConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	return string(d)
}
