package main

import (
	"log"

	"github.com/rchaganti/hello-world/cmd"
)

var (
	version = ""
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.SetVersionInfo(version, commit, date)

	// call the root command's Execute function
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
