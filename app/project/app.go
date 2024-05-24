package project

import (
	"os"
	"sync"

	"github.com/ryanadiputraa/ggen/app/cache"
	"github.com/ryanadiputraa/ggen/config"
)

func writeApp(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.MkdirAll("app/ggen/delivery/http/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("app/ggen/repository/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("app/ggen/service/", userPermission); err != nil {
		return
	}

	wg := sync.WaitGroup{}
	errChan := make(chan error, 4)

	runTask(&wg, errChan, func() (err error) {
		cache.Delivery, err = generateTemplateFile(config, "/app/template/app/ggen/delivery/http/delivery.go", "app/ggen/delivery/http/delivery.go", cache.Delivery, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Repository, err = generateTemplateFile(config, "/app/template/app/ggen/repository/repository.go", "app/ggen/repository/repository.go", cache.Repository, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Service, err = generateTemplateFile(config, "/app/template/app/ggen/service/service.go", "app/ggen/service/service.go", cache.Service, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Ggen, err = generateTemplateFile(config, "/app/template/app/ggen/ggen.go", "app/ggen/ggen.go", cache.Ggen, isUseCache)
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
