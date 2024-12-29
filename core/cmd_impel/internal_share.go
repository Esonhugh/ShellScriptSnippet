package cmd_impel

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Esonhugh/ShellScriptSnippet/core/defines"
	"github.com/Esonhugh/ShellScriptSnippet/utils/log"
	"github.com/guonaihong/gout"
)

const (
	GistApi     = "https://api.github.com/gists"
	PasteBinApi = "https://pastebin.com/api/api_post.php"
)

type GistFile struct {
	Content string `json:"content"`
}

// '{"description":"Example of a gist","public":false,"files":{"README.md":{"content":"Hello World"}}}'
type GistJson struct {
	Description string              `json:"description"`
	Public      bool                `json:"public"`
	Files       map[string]GistFile `json:"files"`
}

type GistResponse struct {
	Url string `json:"html_url"`
}

type PasteJson struct {
	ApiDevKey      string `json:"api_dev_key" form:"api_dev_key"`
	ApiOption      string `json:"api_option" form:"api_option"`
	ApiPasteName   string `json:"api_paste_name" form:"api_paste_name"`
	ApiPasteCode   string `json:"api_paste_code" form:"api_paste_code"`
	ApiPasteFormat string `json:"api_paste_format" form:"api_paste_format"`
}

type ShareFunction func(t defines.ShellSnippet, secret string) (string, error)

var ShareOnGithub ShareFunction = func(t defines.ShellSnippet, s string) (string, error) {
	var gistR GistResponse
	r := gout.POST(GistApi).SetHeader(gout.H{
		"Accept":        "application/vnd.github+json",
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + s,
	}).SetJSON(&GistJson{
		Description: "Shared by ShellScriptSnippet",
		Public:      true,
		Files: map[string]GistFile{
			fmt.Sprintf("%s.sh", strings.ReplaceAll(t.Name, " ", "_")): {t.Content},
		},
	}).BindJSON(&gistR).Debug()

	if log.GlobalLogLevel > log.LogLevelInfo {
		r.Debug(true)
	}

	e := r.Do()
	if e != nil {
		return "", e
	}
	return gistR.Url, nil
}

var ShareOnPastebin ShareFunction = func(t defines.ShellSnippet, s string) (string, error) {
	var str string
	r := gout.POST(PasteBinApi).SetForm(&PasteJson{
		ApiDevKey:      s,
		ApiOption:      "paste",
		ApiPasteCode:   t.Content,
		ApiPasteFormat: "bash",
		ApiPasteName:   t.Name,
	}).BindBody(&str)

	if log.GlobalLogLevel > log.LogLevelInfo {
		r.Debug(true)
	}

	e := r.Do()
	if e != nil {
		return "", e
	}
	if !strings.HasPrefix(str, "http") {
		return str, errors.New("bad request")
	}
	return str, nil
}
