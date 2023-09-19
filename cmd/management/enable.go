package management

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"sss/cmd"
	"sss/core/cmd_impel"
)

func init() {
	cmd.RootCmd.AddCommand(EnableCmd, DisableCmd)
}

var EnableCmd = &cobra.Command{
	Use:     "enable",
	Aliases: []string{"e"},
	Short:   "enable shell script snippets",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cmd_impel.EnableSnippetsFromNames(args)
		} else {
			cmd_impel.EnableSnippetsFromSelect()
		}
		log.Infof("enabled successful")
	},
}

var DisableCmd = &cobra.Command{
	Use:     "disable",
	Aliases: []string{"dis"},
	Short:   "disable shell script snippets",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cmd_impel.DisableSnippetsFromNames(args)
		} else {
			cmd_impel.DisableSnippetsFromSelect()
		}
		log.Infof("disabled successful")
	},
}
