package main

import (
	"github.com/Esonhugh/ShellScriptSnippet/cmd"
	_ "github.com/Esonhugh/ShellScriptSnippet/cmd/initrc"
	_ "github.com/Esonhugh/ShellScriptSnippet/cmd/management"
	_ "github.com/Esonhugh/ShellScriptSnippet/cmd/version" // import sub command as module
)

func init() {
}

func main() {
	cmd.Execute()
}
