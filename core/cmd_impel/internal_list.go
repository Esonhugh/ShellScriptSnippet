package cmd_impel

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"sss/utils/Print"
)

func ListAndPrint() {
	table := Print.Table{
		Header: []string{"name", "is_enabled", "content"},
	}
	rs, err := GetAllSnippets()
	if err != nil {
		log.Errorf("List Snippets Error: %v", err)
		os.Exit(-1)
	}
	for _, v := range rs {
		table.Body = append(table.Body, []string{
			v.Name,
			fmt.Sprintf("%v", v.IsEnable),
			v.MaskedContent(),
		})
	}
	table.Print("ALL Snippets")
}
