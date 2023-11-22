package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate go project",
	Long: `Generate go project with idiomatic go standard layout.

	More about the project layout referrence can be seen here:
	https://github.com/golang-standards/project-layout`,
	Run: generateProject,
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generateProject(cmd *cobra.Command, args []string) {
	// generate standard layout directory
	fmt.Println("Generating projects...")

	// c := exec.Command("mkdir", "project-name")
	// if err := c.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}
