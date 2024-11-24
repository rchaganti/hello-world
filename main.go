package main

import (
	"log"

	"github.com/rchaganti/hello-world/cmd"
)

var (
	version = "v0.0.4"
	commit  = "none"
	date    = "unknown"
)

func main() {
	// Set the version information
	cmd.SetVersionInfo(version, commit, date)

	// call the root command's Execute function
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
