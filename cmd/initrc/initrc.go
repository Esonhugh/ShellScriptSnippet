package initrc

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"sss/cmd"
	"sss/core/cmd_impel"
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
		// if ssl_leaded is any and reload is set, reload
		// if sss_loaded not set but reload not set, load
		// if ssl_loaded set but reload not set, no load
		if os.Getenv("SSS_LOADED") == "true" && !reload {
			// NeverLoad Again
			return
		}
		println("export SSS_LOADED=true;")
		cmd_impel.InitRC()
	},
}

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "output how to use this cli in rc files",
	Run: func(cmd *cobra.Command, args []string) {
		fp, _ := filepath.Abs(os.Args[0])
		log.Infoln("using following command to add into zshrc:")
		log.Infof(`	echo 'eval "$(%v initrc)"' >> ~/.zshrc `, fp)
		log.Infoln("using following command to add into bashrc:")
		log.Infof(`	echo 'eval "$(%v initrc)"' >> ~/.bashrc `, fp)
	},
}
