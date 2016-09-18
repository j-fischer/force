package main

import (
	"fmt"
)

var Version = "dev1" // set to dev1 to prevent the callback URL from hitting a staging environment for sandboxes.

//Dood, what
var cmdVersion = &Command{
	Run:   runVersion,
	Usage: "version",
	Short: "Display current version",
	Long: `
Display current version

Examples:

  force version
`,
}

func init() {
}

func runVersion(cmd *Command, args []string) {
	fmt.Println(Version)
}
