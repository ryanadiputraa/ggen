package project

import (
	"os"
	"sync"

	"github.com/ryanadiputraa/ggen/v2/app/cache"
	"github.com/ryanadiputraa/ggen/v2/config"
)

func writeApp(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.MkdirAll("app/server/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("app/health/delivery/http/", userPermission); err != nil {
		return
	}

	wg := sync.WaitGroup{}
	errChan := make(chan error, 4)

	runTask(&wg, errChan, func() (err error) {
		cache.Server, err = generateTemplateFile(config, "/app/template/app/server/server.go", "app/server/server.go", cache.Server, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Handler, err = generateTemplateFile(config, "/app/template/app/server/handlers.go", "app/server/handlers.go", cache.Handler, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Health, err = generateTemplateFile(config, "/app/template/app/health/health.go", "app/health/health.go", cache.Health, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.HealthHandler, err = generateTemplateFile(config, "/app/template/app/health/delivery/http/http.go", "app/health/delivery/http/http.go", cache.HealthHandler, isUseCache)
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
