package initrc

import (
	"fmt"
	"github.com/Esonhugh/ShellScriptSnippet/cmd"
	"github.com/Esonhugh/ShellScriptSnippet/core/cmd_impel"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var reload bool

func init() {
	InitrcCmd.Flags().BoolVarP(&reload, "reload", "r", false, "allow reload even the load flag is set")
	cmd.RootCmd.AddCommand(InitrcCmd, InstallCmd)
}

var InitrcCmd = &cobra.Command{
	Use:   "initrc",
	Short: "init rc is command to invoke the activated bashrc",
	Run: func(cmd *cobra.Command, args []string) {
		// Always Setting the SSS_RELOAD func
		fp := getCurrentExePath()
		reload_func := fmt.Sprintf(`
function SSS_RELOAD () {
	source <(%v initrc --reload=true)
}
`, fp)
		fmt.Println(reload_func)

		// if ssl_leaded is any and reload is set, reload
		// if sss_loaded not set but reload not set, load
		// if ssl_loaded set but reload not set, no load
		if os.Getenv("SSS_LOADED") == "true" && !reload {
			// NeverLoad Again
			return
		}
		fmt.Println("export SSS_LOADED=true;")
		cmd_impel.InitRC()
	},
}

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "output how to use this cli in rc files",
	Run: func(cmd *cobra.Command, args []string) {
		fp := getCurrentExePath()
		log.Infoln("using following command to add into zshrc:")
		log.Infof(`	echo 'source <(%v initrc)' >> ~/.zshrc `, fp)
		log.Infoln("using following command to add into bashrc:")
		log.Infof(`	echo 'source <(%v initrc)' >> ~/.bashrc `, fp)
	},
}

func getCurrentExePath() string {
	ex, _ := os.Executable()
	return ex
}
