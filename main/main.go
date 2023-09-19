package main

import (
	"sss/cmd"
	_ "sss/cmd/version" // import sub command as module
)

func init() {
}

func main() {
	cmd.Execute()
}
