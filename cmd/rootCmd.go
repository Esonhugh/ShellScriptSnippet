package cmd

import (
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"os"
	"sss/utils/log"
)

var RootCmd = &cobra.Command{
	Use:   "sss",
	Short: "ShellScriptSnippet(abbr.sss) is a kind of commandline tool for managing,sharing your self created shell script to others",
	Long:  "#########################################\n#     _______.     _______.     _______.#\n#    /       |    /       |    /       |#\n#   |   (----`   |   (----`   |   (----`#\n#    \\   \\        \\   \\        \\   \\    #\n#.----)   |   .----)   |   .----)   |   #\n#|_______/    |_______/    |_______/    #\n#########################################",
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		log.Init()
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&log.GlobalLogLevel, "verbose", "info", "设置日志等级 (Set log level) [trace|debug|info|warn|error|fatal|panic]")
	// RootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:  RootCmd,
		Headings: cc.HiGreen + cc.Underline,
		Commands: cc.Cyan + cc.Bold,
		Example:  cc.Italic,
		ExecName: cc.Bold,
		Flags:    cc.Cyan + cc.Bold,
	})
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
