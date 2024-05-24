package project

import (
	"os"
	"sync"

	"github.com/ryanadiputraa/ggen/app/cache"
	"github.com/ryanadiputraa/ggen/config"
)

func writeMiddlewares(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.MkdirAll("app/middleware/", userPermission); err != nil {
		return
	}

	wg := sync.WaitGroup{}
	errChan := make(chan error, 2)

	runTask(&wg, errChan, func() (err error) {
		cache.Cors, err = generateTemplateFile(config, "/app/template/app/middleware/cors.go", "app/middleware/cors.go", cache.Cors, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Throttle, err = generateTemplateFile(config, "/app/template/app/middleware/throttle.go", "app/middleware/throttle.go", cache.Throttle, isUseCache)
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
