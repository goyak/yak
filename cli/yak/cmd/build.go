// Copyright © 2017 Shawn Wang <shawn.wang@easystack.cn>
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

package cmd

import (
	"fmt"

	"github.com/goyak/yak/lib/recipe"
)

var buildCmd = appCmd(build, "build")

func build(r recipe.IRecipeConfig) {
	f, ok := r.(recipe.AtomicRecipeConfig)
	fmt.Println(f, ok)
	if ok {
		fmt.Printf("build: check type (%T) \n", r)
	}
}

func init() {
	RootCmd.AddCommand(buildCmd)
}
