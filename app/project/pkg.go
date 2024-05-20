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
	cache.Postgres, err = generateTemplateFile(config, "/app/template/pkg/db/postgres.go", "pkg/db/postgres.go", cache.Postgres, isUseCache)
	return
}
