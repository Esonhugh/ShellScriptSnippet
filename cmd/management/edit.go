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

var EditCmdOpts struct {
	SetNameAsFileName bool
	SetEnable         bool
	UpdateMode        bool
	Name              string
}

func init() {
	EditCmd.Flags().BoolVarP(&EditCmdOpts.SetEnable, "enable", "e", true, "Set snippet default as Enable")
	EditCmd.Flags().BoolVarP(&EditCmdOpts.SetNameAsFileName, "keep-file-name", "k", false, "Set Name as file Name")
	EditCmd.Flags().StringVarP(&EditCmdOpts.Name, "name", "n", "",
		"Overwrite Snippet Name (pair with from-file=false), If name is same, edit command will update that content")
	cmd.RootCmd.AddCommand(EditCmd)
}

var EditCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e"},
	Short:   "Edit/Update a shell script snippet",
	Run: func(cmd *cobra.Command, args []string) {
		// Edit from Editor
		// Name Ask name
		// Content form Editor result
		// Enable set as EditAsSetEnable
		var realname string = EditCmdOpts.Name
		var realenable bool = EditCmdOpts.SetEnable
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
				if rs.Name == EditCmdOpts.Name {
					realname = rs.Name
					realenable = rs.IsEnable
					found = true
					break
				}
			}
			if !found {
				log.Errorf("User input name %v not found in database", EditCmdOpts.Name)
				os.Exit(-1)
			}
		}
		cmd_impel.AddFromEditor(realname, realenable)
	},
}
