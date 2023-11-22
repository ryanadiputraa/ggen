package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ggen",
	Short: "Go project generator",
	Long: `A CLI for generating go project that use idiomatic go project standard layout

More about the project layout referrence can be seen here:
https://github.com/golang-standards/project-layout`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
