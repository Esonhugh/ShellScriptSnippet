package management

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"sss/cmd"
	"sss/core/cmd_impel"
	"sss/core/defines"
	"sss/utils/Ask"
)

var AddCmdOpts struct {
	FromFile          bool
	SetNameAsFileName bool
	SetEnable         bool
	UpdateMode        bool
	Name              string
}

func init() {
	AddCmd.Flags().BoolVarP(&AddCmdOpts.FromFile, "from-file", "f", false, "Add From File")
	AddCmd.Flags().BoolVarP(&AddCmdOpts.SetEnable, "enable", "e", true, "Set snippet default as Enable")
	AddCmd.Flags().BoolVarP(&AddCmdOpts.UpdateMode, "update", "u", false, "Update mode")
	AddCmd.Flags().BoolVarP(&AddCmdOpts.SetNameAsFileName, "keep-file-name", "k", false, "Set Name as file Name")

	AddCmd.Flags().StringVarP(&AddCmdOpts.Name, "name", "n", "",
		"Overwrite Snippet Name (pair with from-file=false), If name is same, add command will update that content")
	cmd.RootCmd.AddCommand(AddCmd)
}

var AddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a shell script snippet",
	Run: func(cmd *cobra.Command, args []string) {
		if AddCmdOpts.FromFile {
			// Add from file
			// files is args
			// filename is the name
			// content is file data
			// Enable set as AddAsSetEnable
			cmd_impel.AddFromFiles(args, AddCmdOpts.SetEnable, AddCmdOpts.SetNameAsFileName)
		} else {
			// Add from Editor
			// Name Ask name
			// Content form Editor result
			// Enable set as AddAsSetEnable
			if AddCmdOpts.UpdateMode {
				allrs, err := cmd_impel.GetAllSnippets()
				if err != nil {
					log.Errorf("Get all Snippets Error: %v", err)
				}
				res := Ask.ChoiceOneOfSlice(Ask.ConvertToChoiceAbleSlice(allrs)).(defines.ShellSnippet)
				cmd_impel.AddFromEditor(res.Name, res.IsEnable)
				return
			}
			if AddCmdOpts.Name == "" {
				AddCmdOpts.Name = Ask.ForName("Input snippet name")
			}
			cmd_impel.AddFromEditor(AddCmdOpts.Name, AddCmdOpts.SetEnable)
		}
	},
}
