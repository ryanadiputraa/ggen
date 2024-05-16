package project

import (
	"os"

	"github.com/ryanadiputraa/ggen/config"
)

func writePkg(config *config.Config) (err error) {
	if err = os.MkdirAll("pkg/db/", userPermission); err != nil {
		return
	}
	return generateTemplateFile(config, "/app/template/pkg/db/postgres.go", "pkg/db/postgres.go")
}
