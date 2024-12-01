package main

import (
	"log"
	"time"

	"github.com/rchaganti/hello-world/cmd"
)

var (
	version = "v0.0.7"
	commit  = "none"
	date    = "unknown"
)

func main() {
	// Set the version information
	date = time.Now().Format(time.RFC3339)
	cmd.SetVersionInfo(version, commit, date)

	// call the root command's Execute function
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
