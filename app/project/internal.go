package project

import (
	"os"
	"sync"

	"github.com/ryanadiputraa/ggen/app/cache"
	"github.com/ryanadiputraa/ggen/config"
)

func writeInternal(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.MkdirAll("internal/server/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("internal/middleware/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("internal/database/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("internal/logger/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("internal/respwr/", userPermission); err != nil {
		return
	}

	wg := sync.WaitGroup{}
	errChan := make(chan error, 7)

	runTask(&wg, errChan, func() (err error) {
		cache.Server, err = generateTemplateFile(config, "/app/template/internal/server/server.go", "internal/server/server.go", cache.Server, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Handler, err = generateTemplateFile(config, "/app/template/internal/server/handler.go", "internal/server/handler.go", cache.Handler, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Cors, err = generateTemplateFile(config, "/app/template/internal/middleware/cors.go", "internal/middleware/cors.go", cache.Cors, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Throttle, err = generateTemplateFile(config, "/app/template/internal/middleware/throttle.go", "internal/middleware/throttle.go", cache.Throttle, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Database, err = generateTemplateFile(config, "/app/template/internal/database/database.go", "internal/database/database.go", cache.Database, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Logger, err = generateTemplateFile(config, "/app/template/internal/logger/logger.go", "internal/logger/logger.go", cache.Logger, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Respwr, err = generateTemplateFile(config, "/app/template/internal/respwr/respwr.go", "internal/respwr/respwr.go", cache.Respwr, isUseCache)
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
