package initrc

import (
	"fmt"
	"github.com/Esonhugh/ShellScriptSnippet/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	cmd.RootCmd.AddCommand(InitrcCmd, InstallCmd)
}

var InitrcCmd = &cobra.Command{
	Use:   "initrc",
	Short: "init rc is command to invoke the activated bashrc",
	Run: func(cmd *cobra.Command, args []string) {
		// Always Setting the SSS_RELOAD func
		fp := getCurrentExePath()
		// make SSS_RELOAD function definition as a target to determine whether the script is loaded
		// bcs the environment variable is not reliable, it will pass to the subprocess shell
		// so if launch a new shell, the defined environment variable will be same but functions and alias will be lost
		reload_func := fmt.Sprintf(`
declare -f SSS_RELOAD > /dev/null && export SSS_LOADED=true || export SSS_LOADED=false
function SSS_RELOAD () {
	source <(%v load --reload=true)
}
if [[ "$SSS_LOADED" == "false" ]]; then
	source <(%v load)
fi
`, fp, fp)
		fmt.Println(reload_func)
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
