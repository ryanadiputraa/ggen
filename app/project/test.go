package project

import (
	"os"
	"sync"

	"github.com/ryanadiputraa/ggen/v2/app/cache"
	"github.com/ryanadiputraa/ggen/v2/config"
)

func writeTest(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.MkdirAll("test/", userPermission); err != nil {
		return
	}

	wg := sync.WaitGroup{}
	errChan := make(chan error, 2)

	runTask(&wg, errChan, func() (err error) {
		cache.TestSetup, err = generateTemplateFile(config, "/app/template/test/setup.go", "test/setup.go", cache.TestSetup, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.TestHealthcheck, err = generateTemplateFile(config, "/app/template/test/test_healthcheck.go", "test/test_healthcheck.go", cache.TestHealthcheck, isUseCache)
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
