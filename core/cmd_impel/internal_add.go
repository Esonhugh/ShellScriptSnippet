package cmd_impel

import (
	"fmt"
	"github.com/Esonhugh/ShellScriptSnippet/core/defines"
	"github.com/Esonhugh/ShellScriptSnippet/utils/Ask"
	"github.com/Esonhugh/ShellScriptSnippet/utils/File"
	"github.com/Esonhugh/ShellScriptSnippet/utils/TextEditorCall"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

func AddFromFiles(files []string, enable bool, keepFileName bool) {
	for _, file := range files {
		if !File.IsFile(file) {
			log.Errorf("file %v not exists or is dir", file)
			return
		}
		var sn defines.ShellSnippet
		if keepFileName {
			fileName := filepath.Base(file)
			baseName := strings.TrimSuffix(fileName, filepath.Ext(fileName))
			sn.Name = baseName
		} else {
			sn.Name = Ask.ForName(fmt.Sprintf("Input file %v, set snippet name as", file))
		}
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
