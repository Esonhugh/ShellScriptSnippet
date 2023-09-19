package defines

import (
	"fmt"
	"strings"
)

type ShellType string

const (
	Powershell = ShellType("PowerShell")
	Zsh        = ShellType("Zsh")
	Fish       = ShellType("Fish")
	Bash       = ShellType("Bash")
	Other      = ShellType("Other")
)

// One ShellSnippet struct
type ShellSnippet struct {
	// Name of ShellSnippet
	Name string `gorm:"primaryKey"`

	// Raw Content of This ShellSnippet
	Content string

	// Enable this shell snippet if registry sss into .bashrc like rc files.
	IsEnable bool
}

func (s ShellSnippet) UniqName() string {
	return fmt.Sprintf("%v - %v", s.Name, s.IsEnable)
}

func (s ShellSnippet) MaskedContent() string {
	res := s.Content
	if len(res) > 60 {
		res = res[:40] + res[len(res)-15:]
	}
	res = strings.Replace(res, "\n", "\\n", -1)
	return res
}
