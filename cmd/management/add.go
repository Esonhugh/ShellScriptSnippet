package management

import (
	"github.com/spf13/cobra"
	"sss/cmd"
)

func init() {
	cmd.RootCmd.AddCommand(AddingShellSnippetCmd)
}

var AddingShellSnippetCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a shell script snippet",
}
