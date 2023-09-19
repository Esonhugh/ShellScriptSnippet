package main

import (
	log2 "github.com/sirupsen/logrus"
	"sss/utils/TextEditorCall"
	"sss/utils/log"
)

func main() {
	log.GlobalLogLevel = "debug"
	log.Init()
	data := TextEditorCall.TextEdit("vim")
	log2.Info(data)
}
