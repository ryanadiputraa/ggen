package project

import (
	"os"

	"github.com/ryanadiputraa/ggen/config"
)

func writeConfigFile(config *config.Config) (err error) {
	if err = os.Chdir(config.ProjectDir); err != nil {
		return
	}
	defer os.Chdir(config.GgenDir)

	if err = os.Mkdir("config", userPermission); err != nil {
		return
	}

	if err = generateTemplateFile(config, "/app/template/config/config.yml", "config/config.yml"); err != nil {
		return
	}
	return generateTemplateFile(config, "/app/template/config/config.go", "config/config.go")
}
