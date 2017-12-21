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
	"log"
	"path/filepath"

	"gitlab.com/EasyStack/yakety/lib/env"
	"gitlab.com/EasyStack/yakety/lib/index"
	"gitlab.com/EasyStack/yakety/lib/recipe"
)

var installCmd = appCmd(install, "install")

func install(r recipe.IRecipeConfig) {
	cfg := r.GetRecipeConfig()
	if !r.IsRecipe() {
		log.Printf("do yak fetch %s %q\n", cfg.Name, cfg)
		r.Fetch(env.YakRoot())
		r = recipe.LoadRecipe(env.YakRoot(), cfg.Repo)
	}
	if !r.IsInstallable() {
		log.Fatalf("cannot install: %s\n", cfg.Name)
		return
	}
	if !r.Install(doRun) {
		log.Printf("install failed %s\n", cfg.Name)
		return
	}

	if doRun {
		path := filepath.Join(env.YakRoot(), env.LocalIndex)
		idx := index.LoadIndex(path)
		idx.Install(cfg)
		log.Printf("index updated.\n")
	}
}

func init() {
	RootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVarP(&doRun, "dry-run", "D", false, "Dry Run")
}
