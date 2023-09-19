package cmd_impel

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"sss/core/defines"
	"sss/utils/Ask"
	"sss/utils/File"
	"sss/utils/TextEditorCall"
)

func AddFromFiles(files []string, enable bool) {
	for _, file := range files {
		if !File.IsFile(file) {
			log.Errorf("file %v not exists or is dir", file)
			return
		}
		var sn defines.ShellSnippet
		sn.Name = Ask.ForName(fmt.Sprintf("Input file %v, set snippet name as", file))
		byteData, err := os.ReadFile(file)
		if err != nil {
			log.Errorf("Read File %v content error: %v", file, err)
			return
		}
		sn.Content = string(byteData)
		sn.IsEnable = enable
		SaveSnippet(sn)
	}
}

func AddFromEditor(name string, enable bool) {
	sn, e := FoundSnippet(name)
	if e != nil {
		// Not Found or DB error
		sn.Content = TextEditorCall.TextEditSurvey("")
		sn.Name = name
		sn.IsEnable = enable
	} else {
		// Duplicated name edit the sn
		sn.Content = TextEditorCall.TextEditSurvey(sn.Content)
	}
	SaveSnippet(sn)
}
