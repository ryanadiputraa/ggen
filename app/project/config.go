package project

import (
	"os"

	"github.com/ryanadiputraa/ggen/app/cache"
	"github.com/ryanadiputraa/ggen/config"
)

func writeConfigFile(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.Mkdir("config", userPermission); err != nil {
		return
	}

	cache.ConfigYML, err = generateTemplateFile(config, "/app/template/config/config.yml", "config/config.yml", cache.ConfigYML, isUseCache)
	if err != nil {
		return
	}
	cache.Config, err = generateTemplateFile(config, "/app/template/config/config.go", "config/config.go", cache.Config, isUseCache)
	return
}
