package main

import (
	"log"

	"github.com/rchaganti/hello-world/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
