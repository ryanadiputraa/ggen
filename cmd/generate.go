package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/ryanadiputraa/ggen/configs"
	"github.com/spf13/cobra"
)

var config *configs.Config

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate go project",
	Long: `Generate go project with idiomatic go standard layout.

	More about the project layout referrence can be seen here:
	https://github.com/golang-standards/project-layout`,
	Run: generateProject,
}

func init() {
	var err error
	config, err = configs.LoadConfig("yml", "configs/config.yml")
	fmt.Println("t: ", config.Name)
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("name", "n", config.Name, "Init project name")
	generateCmd.Flags().StringP("mod", "m", config.Mod, "Go mod name")
}

func generateProject(cmd *cobra.Command, args []string) {
	// generate standard layout directory
	fmt.Println("Generating projects...")
	name, _ := cmd.Flags().GetString("name")
	mod, _ := cmd.Flags().GetString("mod")
	fmt.Println(mod)

	c := exec.Command("mkdir", name)
	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}
