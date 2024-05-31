package project

import (
	"os"
	"sync"

	"github.com/ryanadiputraa/ggen/v2/app/cache"
	"github.com/ryanadiputraa/ggen/v2/config"
)

func writeMain(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.MkdirAll("cmd/api/", userPermission); err != nil {
		return
	}

	wg := sync.WaitGroup{}
	errChan := make(chan error, 3)

	runTask(&wg, errChan, func() (err error) {
		cache.CMD, err = generateTemplateFile(config, "/app/template/cmd/api/main.go", "cmd/api/main.go", cache.CMD, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Makefile, err = generateTemplateFile(config, "/app/template/Makefile", "Makefile", cache.Makefile, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Readme, err = generateTemplateFile(config, "/app/template/README.md", "README.md", cache.Readme, isUseCache)
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
