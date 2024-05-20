package project

import (
	"os"

	"github.com/ryanadiputraa/ggen/app/cache"
	"github.com/ryanadiputraa/ggen/config"
)

func writeApp(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.MkdirAll("app/ggen/delivery/http/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("app/ggen/repository/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("app/ggen/service/", userPermission); err != nil {
		return
	}

	cache.Delivery, err = generateTemplateFile(config, "/app/template/app/ggen/delivery/http/delivery.go", "app/ggen/delivery/http/delivery.go", cache.Delivery, isUseCache)
	if err != nil {
		return
	}
	cache.Repository, err = generateTemplateFile(config, "/app/template/app/ggen/repository/repository.go", "app/ggen/repository/repository.go", cache.Repository, isUseCache)
	if err != nil {
		return
	}
	cache.Service, err = generateTemplateFile(config, "/app/template/app/ggen/service/service.go", "app/ggen/service/service.go", cache.Service, isUseCache)
	if err != nil {
		return
	}
	cache.Ggen, err = generateTemplateFile(config, "/app/template/app/ggen/ggen.go", "app/ggen/ggen.go", cache.Ggen, isUseCache)
	return
}
