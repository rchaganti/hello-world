package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hello-world",
	Short: "A CLI demo application for Go",
	Long: `A CLI demo application written in Go to demonstrate
	       distributing CLI applications and tools. 
		   
		   This application can be installed using the go install command.`,
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		// get the first arg
		if len(args) == 0 {
			name = "World"
		} else {
			name = args[0]
		}
		fmt.Printf("Hello, %s. Welcome to the Go world!\n", name)
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

func SetVersionInfo(version, commit, date string) {
	rootCmd.Version = fmt.Sprintf("%s (Built on %s from Git SHA %s)", version, date, commit)
}
