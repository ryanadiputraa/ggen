package project

import (
	"os"

	"github.com/ryanadiputraa/ggen/v2/app/cache"
	"github.com/ryanadiputraa/ggen/v2/config"
)

func writeCMD(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.MkdirAll("cmd/api/", userPermission); err != nil {
		return
	}
	cache.CMD, err = generateTemplateFile(config, "/app/template/cmd/api/main.go", "cmd/api/main.go", cache.CMD, isUseCache)
	if err != nil {
		return
	}
	return
}
