package management

import (
	"github.com/Esonhugh/ShellScriptSnippet/cmd"
	"github.com/Esonhugh/ShellScriptSnippet/core/cmd_impel"
	"github.com/Esonhugh/ShellScriptSnippet/core/defines"
	"github.com/Esonhugh/ShellScriptSnippet/utils/Ask"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
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
	Short:   "Add/Update a shell script snippet",
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
			var realname string = AddCmdOpts.Name
			var realenable bool = AddCmdOpts.SetEnable
			if AddCmdOpts.UpdateMode {
				allrs, err := cmd_impel.GetAllSnippets()
				if err != nil {
					log.Errorf("Get all Snippets Error: %v", err)
				}
				if realname == "" {
					res := Ask.ChoiceOneOfSlice(Ask.ConvertToChoiceAbleSlice(allrs)).(defines.ShellSnippet)
					realname = res.Name
					realenable = res.IsEnable
				} else {
					found := false
					for _, rs := range allrs {
						if rs.Name == AddCmdOpts.Name {
							realname = rs.Name
							realenable = rs.IsEnable
							found = true
							break
						}
					}
					if !found {
						log.Errorf("User input name %v not found in database", AddCmdOpts.Name)
						os.Exit(-1)
					}
				}
			} else {
				if realname == "" {
					realname = Ask.ForName("Input snippet name")
				}
			}
			cmd_impel.AddFromEditor(realname, realenable)
		}
	},
}
