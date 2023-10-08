package management

import (
	"github.com/Esonhugh/ShellScriptSnippet/cmd"
	"github.com/Esonhugh/ShellScriptSnippet/core/cmd_impel"
	"github.com/Esonhugh/ShellScriptSnippet/core/defines"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	ShareCmd.Flags().BoolVarP(&ShareCmdOpts.DryRun, "dry-run", "d", true, "Print only not put online")
	ShareCmd.Flags().StringVarP(&ShareCmdOpts.Name, "name", "n", "", "Snippet name")
	cmd.RootCmd.AddCommand(ShareCmd)
}

var ShareCmdOpts struct {
	Name   string
	DryRun bool
}

var ShareCmd = &cobra.Command{
	Use:     "share",
	Aliases: []string{"s"},
	Short:   "Share or print snippets",
	Run: func(cmd *cobra.Command, args []string) {
		var target defines.ShellSnippet
		if ShareCmdOpts.Name == "" {
			target = cmd_impel.SelectSnippet()
		} else {
			t, err := cmd_impel.FoundSnippet(ShareCmdOpts.Name)
			if err != nil {
				log.Errorf("Get Share Snippet error: %v", err)
				os.Exit(-1)
			}
			target = t
		}
		if ShareCmdOpts.DryRun {
			cmd_impel.PrintTarget(target)
		} else {
			// cmd_impel.ShareTarget(platform, target)
		}
	},
}
