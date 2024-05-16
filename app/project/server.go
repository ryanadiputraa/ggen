package project

import (
	"os"

	"github.com/ryanadiputraa/ggen/config"
)

func writeServer(config *config.Config) (err error) {
	if err = os.MkdirAll("app/server/", userPermission); err != nil {
		return
	}

	if err = generateTemplateFile(config, "/app/template/app/server/server.go", "app/server/server.go"); err != nil {
		return
	}
	return generateTemplateFile(config, "/app/template/app/server/handler.go", "app/server/handler.go")
}
