package main

//
// File is for the test that can't run in goland test

import (
	"github.com/Esonhugh/ShellScriptSnippet/utils/Ask"
	"github.com/Esonhugh/ShellScriptSnippet/utils/TextEditorCall"
	"github.com/Esonhugh/ShellScriptSnippet/utils/log"
	log2 "github.com/sirupsen/logrus"
)

func main() {
	log.GlobalLogLevel = "debug"
	log.Init()
	// main3()
	main1_5()
}

func main1() {
	data := TextEditorCall.TextEdit("vim")
	log2.Info(data)
}

func main1_5() {
	data := TextEditorCall.TextEditSurvey("default")
	log2.Info(data)
}

type TestChoice struct {
	Name string
}

func (t TestChoice) UniqName() string {
	return t.Name
}

func main2() {
	var data = []Ask.ChoiceAble{
		TestChoice{Name: "a"},
		TestChoice{Name: "b"},
		TestChoice{Name: "c"},
		TestChoice{Name: "d"},
		TestChoice{Name: "e"},
		TestChoice{Name: "f"},
		TestChoice{Name: "g"},
		TestChoice{Name: "h"},
	}

	result := Ask.ChoiceMultiOfSlice(data)
	for _, v := range result {
		println("Choice: ", v.UniqName())
	}
}

func main3() {
	var data = []TestChoice{
		TestChoice{Name: "a"},
		TestChoice{Name: "b"},
		TestChoice{Name: "c"},
		TestChoice{Name: "d"},
		TestChoice{Name: "e"},
		TestChoice{Name: "f"},
		TestChoice{Name: "g"},
		TestChoice{Name: "h"},
	}
	result := Ask.ChoiceOneOfSlice(Ask.ConvertToChoiceAbleSlice(data))
	println(result.(TestChoice).Name)
}
