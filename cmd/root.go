package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Version: "v1.3.1",
	Use:     "ggen",
	Short:   "Go project generator",
	Long: `A CLI for generating go project, it helps automate the process of creating a new Go project with a predefined directory structure, configuration files, and third party library/package.

More about the project layout referrence can be seen here:
https://github.com/golang-standards/project-layout`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
