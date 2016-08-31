package main

import "github.com/spf13/cobra/doc"
import "github.com/FederationOfFathers/destiny/destiny-cli/cmd"

func main() {
	header := &doc.GenManHeader{
		Title:   "vgo",
		Section: "3",
	}
	doc.GenManTree(cmd.RootCmd, header, "./")
}
