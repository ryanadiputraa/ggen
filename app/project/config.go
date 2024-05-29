package project

import (
	"os"
	"sync"

	"github.com/ryanadiputraa/ggen/v2/app/cache"
	"github.com/ryanadiputraa/ggen/v2/config"
)

func writeConfigFile(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.Mkdir("config", userPermission); err != nil {
		return
	}

	wg := sync.WaitGroup{}
	errChan := make(chan error, 2)

	runTask(&wg, errChan, func() (err error) {
		cache.ConfigYML, err = generateTemplateFile(config, "/app/template/config/config.yml", "config/config.yml", cache.ConfigYML, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Config, err = generateTemplateFile(config, "/app/template/config/config.go", "config/config.go", cache.Config, isUseCache)
		return
	})

	wg.Wait()
	close(errChan)

	for e := range errChan {
		if e != nil {
			err = e
		}
	}

	return
}
