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
	"path/filepath"

	"gitlab.com/EasyStack/yakety/lib/env"
	"gitlab.com/EasyStack/yakety/lib/index"
	"gitlab.com/EasyStack/yakety/lib/recipe"
)

var installCmd = appCmd(install, "install")

func install(r recipe.IRecipeConfig) {
	cfg := r.GetRecipeConfig()
	if !r.IsRecipe() {
		fmt.Printf("need: yak fetch %s\n", cfg.Name)
		return
	}
	if !r.Install() {
		fmt.Printf("install failed %s\n", cfg.Name)
		return
	}
	path := filepath.Join(env.YakRoot(), env.LocalIndex)
	idx := index.LoadIndex(path)
	idx.Save()
	// TODO add list note
	fmt.Printf("installed %s\n", cfg.Name)
}

func init() {
	RootCmd.AddCommand(installCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	installCmd.Flags().BoolP("toggle_abc", "t", false, "Help message for toggle")
}
