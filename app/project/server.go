package project

import (
	"os"

	"github.com/ryanadiputraa/ggen/app/cache"
	"github.com/ryanadiputraa/ggen/config"
)

func writeServer(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.MkdirAll("app/server/", userPermission); err != nil {
		return
	}

	cache.Server, err = generateTemplateFile(config, "/app/template/app/server/server.go", "app/server/server.go", cache.Server, isUseCache)
	if err != nil {
		return
	}
	cache.Handler, err = generateTemplateFile(config, "/app/template/app/server/handler.go", "app/server/handler.go", cache.Handler, isUseCache)
	return
}
