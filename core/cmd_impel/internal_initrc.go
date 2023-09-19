package cmd_impel

import (
	"fmt"
	"sss/core/defines"
)

func InitRC() {
	DisableGormDebugging()
	rs, e := GetAllEnabledSnippet()
	if e != nil {
		return
	}
	for _, v := range rs {
		println("## SSS_CUSTOM_SNIPPET: " + v.Name)
		println(v.Content)
	}
}

func PrintTarget(r defines.ShellSnippet) {
	println("## SSS_CUSTOM_SNIPPET: " + r.Name)
	println(fmt.Sprintf("## Status: %v", r.IsEnable))
	println(r.Content)
}
