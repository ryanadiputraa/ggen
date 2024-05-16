package project

import (
	"os"

	"github.com/ryanadiputraa/ggen/config"
)

func writeCMD(config *config.Config) (err error) {
	if err = os.MkdirAll("cmd/api/", userPermission); err != nil {
		return
	}
	return generateTemplateFile(config, "/app/template/cmd/api/main.go", "cmd/api/main.go")
}
