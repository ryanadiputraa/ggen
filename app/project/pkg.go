package project

import (
	"os"

	"github.com/ryanadiputraa/ggen/app/cache"
	"github.com/ryanadiputraa/ggen/config"
)

func writePkg(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.MkdirAll("pkg/db/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("pkg/respwr/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("pkg/logger/", userPermission); err != nil {
		return
	}

	cache.Postgres, err = generateTemplateFile(config, "/app/template/pkg/db/postgres.go", "pkg/db/postgres.go", cache.Postgres, isUseCache)
	if err != nil {
		return
	}
	cache.Respwr, err = generateTemplateFile(config, "/app/template/pkg/respwr/respwr.go", "pkg/respwr/respwr.go", cache.Respwr, isUseCache)
	if err != nil {
		return
	}
	cache.Logger, err = generateTemplateFile(config, "/app/template/pkg/logger/logger.go", "pkg/logger/logger.go", cache.Logger, isUseCache)
	return
}
