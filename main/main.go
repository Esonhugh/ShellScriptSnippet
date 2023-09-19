package main

import (
	"sss/cmd"
	_ "sss/cmd/initrc"
	_ "sss/cmd/management"
	_ "sss/cmd/version" // import sub command as module
)

func init() {
}

func main() {
	cmd.Execute()
}
