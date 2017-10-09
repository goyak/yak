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

	"gopkg.in/yaml.v2"
)

type RecipeConfig struct {
	Name        string `yaml:"name"`
	Summary     string
	Version     string
	Backend     string
	Source      string
	Branch      string
	Hash        string
	Description string
	Configs     struct {
		Config map[string]string `yaml:",inline"`
	} `yaml:"config"`
}

type Recipe struct {
	RecipeConfig
}

type IRecipe interface {
	IsInstallable() bool
	Dump() string
}

func LoadRecipe(file string) IRecipe {
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
		r := AtomicRecipe{}
		r.RecipeConfig = config
		return r
	} else {
		r := Recipe{}
		r.RecipeConfig = config
		return r
	}
}

func (r Recipe) IsInstallable() bool {
	return true
}

func (r Recipe) Dump() string {
	d, err := yaml.Marshal(&r.RecipeConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	return string(d)
}
