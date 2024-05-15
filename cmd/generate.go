package cmd

import (
	"fmt"
	"log"

	"github.com/ryanadiputraa/ggen/app/module"
	"github.com/ryanadiputraa/ggen/config"
	"github.com/spf13/cobra"
)

const (
	DefaultName = "go-project"
	DefaultMod  = "github.com/username/go-project"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate go project",
	Long: `Generate go project with idiomatic go standard layout.

Usage example:
ggen generate -n ggen -m github.com/ryanadiputraa/ggen

More about the project layout referrence can be seen here:
https://github.com/golang-standards/project-layout`,
	Run: generateProject,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("name", "n", DefaultName, "Init project name")
	generateCmd.Flags().StringP("mod", "m", DefaultMod, "Go mod name")
}

func generateProject(cmd *cobra.Command, args []string) {
	fmt.Println("Generating projects...")
	name, mod, err := getFlags(cmd)
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.NewConfig(name, mod)
	if err = module.NewModule(cfg); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Project generated!")
}

func getFlags(cmd *cobra.Command) (name, mod string, err error) {
	name, err = cmd.Flags().GetString("name")
	if err != nil {
		return
	}
	mod, err = cmd.Flags().GetString("mod")
	return
}
