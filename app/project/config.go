package project

import (
	"os"
	"sync"

	"github.com/ryanadiputraa/ggen/v2/app/cache"
	"github.com/ryanadiputraa/ggen/v2/config"
)

func writeConfig(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.Mkdir("config", userPermission); err != nil {
		return
	}

	wg := sync.WaitGroup{}
	errChan := make(chan error, 3)

	runTask(&wg, errChan, func() (err error) {
		cache.Env, err = generateTemplateFile(config, "/app/template/.env.example", ".env.example", cache.Env, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Compose, err = generateTemplateFile(config, "/app/template/docker-compose.yml", "docker-compose.yml", cache.Compose, isUseCache)
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
