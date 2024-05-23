package project

import (
	"os"

	"github.com/ryanadiputraa/ggen/app/cache"
	"github.com/ryanadiputraa/ggen/config"
)

func writeMiddlewares(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.MkdirAll("app/middleware/", userPermission); err != nil {
		return
	}

	cache.Cors, err = generateTemplateFile(config, "/app/template/app/middleware/cors.go", "app/middleware/cors.go", cache.Cors, isUseCache)
	if err != nil {
		return
	}
	cache.Throttle, err = generateTemplateFile(config, "/app/template/app/middleware/throttle.go", "app/middleware/throttle.go", cache.Throttle, isUseCache)
	return
}
