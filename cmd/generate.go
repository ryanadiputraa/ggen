package cmd

import (
	"fmt"
	"log"
	"os/exec"

	_cmd "github.com/ryanadiputraa/ggen/internal/cmd"
	"github.com/ryanadiputraa/ggen/internal/configs"
	"github.com/ryanadiputraa/ggen/internal/logger"
	_mod "github.com/ryanadiputraa/ggen/internal/mod"
	"github.com/ryanadiputraa/ggen/internal/server"
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
	name, _ := cmd.Flags().GetString("name")
	mod, _ := cmd.Flags().GetString("mod")
	fmt.Println(mod)

	c := exec.Command("mkdir", name)
	if err := c.Run(); err != nil {
		log.Fatal(err)
	}

	if err := _mod.Write(mod, name); err != nil {
		log.Fatal()
	}
	if err := _cmd.Write(mod, name); err != nil {
		log.Fatal()
	}
	if err := configs.Write(mod, name); err != nil {
		log.Fatal()
	}
	if err := server.Write(mod, name); err != nil {
		log.Fatal()
	}
	if err := logger.Write(mod, name); err != nil {
		log.Fatal()
	}
}
