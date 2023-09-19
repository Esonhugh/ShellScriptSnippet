package management

import (
	"github.com/spf13/cobra"
	"sss/cmd"
	"sss/core/cmd_impel"
)

func init() {
	cmd.RootCmd.AddCommand(ListCmd)
}

var ListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "List and print snippets summary as a table",
	Run: func(cmd *cobra.Command, args []string) {
		cmd_impel.ListAndPrint()
	},
}
