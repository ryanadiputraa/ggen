package project

import (
	"os"

	"github.com/ryanadiputraa/ggen/config"
)

func writeApp(config *config.Config) (err error) {
	if err = os.MkdirAll("app/ggen/delivery/http/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("app/ggen/repository/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("app/ggen/service/", userPermission); err != nil {
		return
	}

	if err = generateTemplateFile(config, "/app/template/app/ggen/delivery/http/delivery.go", "app/ggen/delivery/http/delivery.go"); err != nil {
		return
	}
	if err = generateTemplateFile(config, "/app/template/app/ggen/repository/repository.go", "app/ggen/repository/repository.go"); err != nil {
		return
	}
	if err = generateTemplateFile(config, "/app/template/app/ggen/service/service.go", "app/ggen/service/service.go"); err != nil {
		return
	}
	return generateTemplateFile(config, "/app/template/app/ggen/ggen.go", "app/ggen/ggen.go")
}
