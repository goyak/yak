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
	// "path/filepath"

	"github.com/spf13/cobra"
	"gitlab.com/EasyStack/yakety/lib/host/ostree"
	// "gitlab.com/EasyStack/yakety/lib/index"
)

// listCmd represents the list command
var doList bool

//var idx index.Index
var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Revert to the previously booted tree",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var backup string
		fmt.Printf("rollback called %d, %v\n", len(args), doList)
		if doList {
			fmt.Printf("prepare backlist\n")
			for _, v := range ostree.GetBackupList().Backups {
				fmt.Println(v.Checksum)
			}
			return
		}
		if len(args) == 1 {
			backup = args[0]
		}
		x := ostree.GetRollbackDeployment(backup)
		fmt.Printf("do backup %v\n", x)
	},
}

func init() {
	if ostree.IsOstreeHost() {
		RootCmd.AddCommand(rollbackCmd)
	}
	rollbackCmd.Flags().BoolVarP(&doList, "list", "l", false, "List rollback items")
}
