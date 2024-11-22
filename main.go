package main

import (
	"log"

	"github.com/rchaganti/hello-world/cmd"
)

func main() {
	// call the root command's Execute function
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
