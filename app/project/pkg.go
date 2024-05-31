package project

import (
	"os"
	"sync"

	"github.com/ryanadiputraa/ggen/v2/app/cache"
	"github.com/ryanadiputraa/ggen/v2/config"
)

func writePkg(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.MkdirAll("pkg/db/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("pkg/logger/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("pkg/middleware/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("pkg/respwr/", userPermission); err != nil {
		return
	}

	wg := sync.WaitGroup{}
	errChan := make(chan error, 5)

	runTask(&wg, errChan, func() (err error) {
		cache.Postgres, err = generateTemplateFile(config, "/app/template/pkg/db/postgres.go", "pkg/db/postgres.go", cache.Postgres, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Logger, err = generateTemplateFile(config, "/app/template/pkg/logger/logger.go", "pkg/logger/logger.go", cache.Logger, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Cors, err = generateTemplateFile(config, "/app/template/pkg/middleware/cors.go", "pkg/middleware/cors.go", cache.Cors, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Throttle, err = generateTemplateFile(config, "/app/template/pkg/middleware/throttle.go", "pkg/middleware/throttle.go", cache.Throttle, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Respwr, err = generateTemplateFile(config, "/app/template/pkg/respwr/http-respwr.go", "pkg/respwr/http-respwr.go", cache.Respwr, isUseCache)
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
