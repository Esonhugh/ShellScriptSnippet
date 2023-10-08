package management

import (
	"github.com/Esonhugh/ShellScriptSnippet/cmd"
	"github.com/Esonhugh/ShellScriptSnippet/core/cmd_impel"
	"github.com/spf13/cobra"
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
