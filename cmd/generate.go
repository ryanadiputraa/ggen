package cmd

import (
	"fmt"

	"github.com/ryanadiputraa/ggen/v2/app/module"
	"github.com/ryanadiputraa/ggen/v2/config"
	"github.com/ryanadiputraa/ggen/v2/pkg/logger"
	"github.com/spf13/cobra"
)

const (
	DefaultName = "go-project"
	DefaultMod  = "github.com/username/go-project"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate go project",
	Long:    "Generate go project with a predefined directory structure, configuration files, and third party library/package.",
	Example: "ggen generate -n ggen -m github.com/ryanadiputraa/ggen",
	Run:     generateProject,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("name", "n", DefaultName, "Project name")
	generateCmd.Flags().StringP("mod", "m", DefaultMod, "Go mod name")
}

func generateProject(cmd *cobra.Command, args []string) {
	log := logger.NewLogger()

	name, mod, err := getFlags(cmd)
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.NewConfig(name, mod)
	log.Info(fmt.Sprintf("Generating %v...", cfg.ProjectName))

	// Generate project
	if err = module.NewModule(cfg); err != nil {
		log.Fatal(err)
	}
	if err = module.TidyGoMod(); err != nil {
		log.Fatal(err)
	}
	log.Info("Project generated!")
}

func getFlags(cmd *cobra.Command) (name, mod string, err error) {
	name, err = cmd.Flags().GetString("name")
	if err != nil {
		return
	}
	mod, err = cmd.Flags().GetString("mod")
	return
}
