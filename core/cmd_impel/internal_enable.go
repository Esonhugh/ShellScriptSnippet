package cmd_impel

import (
	log "github.com/sirupsen/logrus"
	"os"
	"sss/core/defines"
	"sss/utils/Database"
)

func EnableSnippetsFromNames(names []string) {
	res := Database.DB.Model(&defines.ShellSnippet{}).Where("name IN ?", names).Update("is_enable", true)
	if res.Error != nil {
		log.Errorf("getting snippets via name error: %v", res.Error)
		os.Exit(-1)
	}
}

func DisableSnippetsFromNames(names []string) {
	res := Database.DB.Model(&defines.ShellSnippet{}).Where("name IN ?", names).Update("is_enable", false)
	if res.Error != nil {
		log.Errorf("getting snippets via name error: %v", res.Error)
		os.Exit(-1)
	}
}

func EnableSnippetsFromSelect() {
	EnableSnippetsFromNames(SelectSnippetsReturnStringSlice())
}

func DisableSnippetsFromSelect() {
	DisableSnippetsFromNames(SelectSnippetsReturnStringSlice())
}
