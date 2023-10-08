package cmd_impel

import (
	"fmt"
	"github.com/Esonhugh/ShellScriptSnippet/core/defines"
)

func InitRC() {
	DisableGormDebugging()
	rs, e := GetAllEnabledSnippet()
	if e != nil {
		return
	}
	for _, v := range rs {
		fmt.Println("## SSS_CUSTOM_SNIPPET: " + v.Name)
		fmt.Println(v.Content)
	}
}

func PrintTarget(r defines.ShellSnippet) {
	fmt.Println("## SSS_CUSTOM_SNIPPET: " + r.Name)
	fmt.Println(fmt.Sprintf("## Status: %v", r.IsEnable))
	fmt.Println(r.Content)
}
