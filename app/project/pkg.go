package project

import (
	"os"
	"sync"

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

	wg := sync.WaitGroup{}
	errChan := make(chan error, 3)

	runTask(&wg, errChan, func() (err error) {
		cache.Postgres, err = generateTemplateFile(config, "/app/template/pkg/db/postgres.go", "pkg/db/postgres.go", cache.Postgres, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Respwr, err = generateTemplateFile(config, "/app/template/pkg/respwr/respwr.go", "pkg/respwr/respwr.go", cache.Respwr, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Logger, err = generateTemplateFile(config, "/app/template/pkg/logger/logger.go", "pkg/logger/logger.go", cache.Logger, isUseCache)
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
