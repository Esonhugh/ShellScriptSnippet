package cmd_impel

import (
	log "github.com/sirupsen/logrus"
	"sss/core/defines"
	"sss/utils/Database"
)

func DeleteSnippetsFromNames(names []string) {
	var snippets []defines.ShellSnippet
	for _, n := range names {
		snippets = append(snippets, defines.ShellSnippet{Name: n})
	}
	deleteSnippets(snippets)
}

func deleteSnippets(rs []defines.ShellSnippet) {
	res := Database.DB.Delete(&rs)
	if res.Error != nil {
		log.Errorf("Deleting snippets meets error: %v", res.Error)
	}
	return
}

func DeleteSnippetsFromSelect() {
	deleteSnippets(SelectSnippets())
}
