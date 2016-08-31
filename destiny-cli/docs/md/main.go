package main

import (
	"github.com/FederationOfFathers/destiny/destiny-cli/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	doc.GenMarkdownTree(cmd.RootCmd, "./")
}
