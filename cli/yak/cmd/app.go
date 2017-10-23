package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/EasyStack/yakety/lib/recipe"
)

func appCmd(fn func(recipe.IRecipeConfig), str string) *cobra.Command {
	return &cobra.Command{
		Use:   str + " [app]",
		Short: "list all installed recipes",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("info Print: " + strings.Join(args, " "))
			fmt.Println("app: " + args[0])
			r := recipe.LoadRecipe(YakRoot(), args[0])
			if r != nil {
				fn(r)
			}
		},
	}
}
