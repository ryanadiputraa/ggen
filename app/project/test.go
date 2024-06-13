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
		cache.SetupTest, err = generateTemplateFile(config, "/app/template/test/setup.go", "test/setup.go", cache.SetupTest, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.HealthcheckTest, err = generateTemplateFile(config, "/app/template/test/healthcheck_test.go", "test/healthcheck_test.go", cache.HealthcheckTest, isUseCache)
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
