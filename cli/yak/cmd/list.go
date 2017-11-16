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

	"github.com/spf13/cobra"
	"gitlab.com/EasyStack/yakety/lib/env"
	"gitlab.com/EasyStack/yakety/lib/index"
)

// listCmd represents the list command
var idx index.Index
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			idx = index.LoadIndex(env.YakRoot() + "/index/local.yml")
		}
		for i, app := range idx.Apps {
			fmt.Printf("app(%d): %v \n", i, app)
		}
		fmt.Printf("list called %d \n", len(args))
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
