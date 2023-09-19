package management

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"sss/cmd"
	"sss/core/cmd_impel"
)

func init() {
	cmd.RootCmd.AddCommand(DelCmd)
}

var DelCmd = &cobra.Command{
	Use:     "del",
	Aliases: []string{"d", "rm", "remove", "r"},
	Short:   "Delete shell script snippets",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cmd_impel.DeleteSnippetsFromNames(args)
		} else {
			cmd_impel.DeleteSnippetsFromSelect()
		}
		log.Infof("delete successful")
	},
}
