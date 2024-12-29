package management

import (
	"errors"
	"fmt"
	"os"

	"github.com/Esonhugh/ShellScriptSnippet/cmd"
	"github.com/Esonhugh/ShellScriptSnippet/core/cmd_impel"
	"github.com/Esonhugh/ShellScriptSnippet/core/defines"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	ShareCmd.Flags().BoolVarP(&ShareCmdOpts.DryRun, "dry-run", "d", false, "Print only not put online")
	ShareCmd.Flags().StringVarP(&ShareCmdOpts.Name, "name", "n", "", "Snippet name")
	ShareCmd.Flags().StringVarP(&ShareCmdOpts.Platform, "platform", "p", "", "Share to platform, supported 'github' 'pastebin' ")
	ShareCmd.Flags().StringVarP(&ShareCmdOpts.PlatformSecret, "secret", "s", "not set", "platform api token")
	cmd.RootCmd.AddCommand(ShareCmd)
}

var ShareCmdOpts struct {
	Name           string
	DryRun         bool
	Platform       string
	PlatformSecret string
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
		if ShareCmdOpts.DryRun || ShareCmdOpts.Platform == "" {
			cmd_impel.PrintTarget(target)
		} else {
			if ShareCmdOpts.PlatformSecret == "not set" {
				log.Errorf("Platform secret not set")
				os.Exit(-1)
			}
			var (
				e error
				u string = ""
			)
			// cmd_impel.ShareTarget(platform, target)
			switch ShareCmdOpts.Platform {
			case "github":
				u, e = cmd_impel.ShareOnGithub(target, ShareCmdOpts.PlatformSecret)
			case "pastebin":
				u, e = cmd_impel.ShareOnPastebin(target, ShareCmdOpts.PlatformSecret)
			default:
				e = errors.New(fmt.Sprintf("Platform %s not support", ShareCmdOpts.Platform))
			}
			if e != nil {
				log.Errorf("Share Error: %v", e)
			} else {
				log.Infof("Share success on url: %v", u)
			}
		}
	},
}
