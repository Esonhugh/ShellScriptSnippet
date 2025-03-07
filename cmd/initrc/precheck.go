package initrc

import (
	"github.com/Esonhugh/ShellScriptSnippet/cmd"
	"github.com/Esonhugh/ShellScriptSnippet/core/cmd_impel"
	"github.com/spf13/cobra"
	"os"
)

var reload bool

func init() {
	LoadShellscriptsCmd.Flags().BoolVarP(&reload, "reload", "r", false, "allow reload even the load flag is set")
	cmd.RootCmd.AddCommand(LoadShellscriptsCmd)
}

var LoadShellscriptsCmd = &cobra.Command{
	Use:   "load",
	Short: "load shell scripts",
	Run: func(cmd *cobra.Command, args []string) {
		// if ssl_leaded is any and reload is set, reload
		// if sss_loaded not set but reload not set, load
		// if ssl_loaded set but reload not set, no load
		if os.Getenv("SSS_LOADED") == "true" && !reload {
			// NeverLoad Again
			return
		}
		cmd_impel.InitRC()
	},
}
