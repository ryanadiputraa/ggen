package project

import (
	"os"
	"sync"

	"github.com/ryanadiputraa/ggen/app/cache"
	"github.com/ryanadiputraa/ggen/config"
)

func writeApp(config *config.Config, isUseCache bool, cache *cache.Cache) (err error) {
	if err = os.MkdirAll("app/template/delivery/http/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("app/template/repository/", userPermission); err != nil {
		return
	}
	if err = os.MkdirAll("app/template/service/", userPermission); err != nil {
		return
	}

	wg := sync.WaitGroup{}
	errChan := make(chan error, 4)

	runTask(&wg, errChan, func() (err error) {
		cache.Delivery, err = generateTemplateFile(config, "/app/template/app/template/delivery/http/delivery.go", "app/template/delivery/http/delivery.go", cache.Delivery, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Repository, err = generateTemplateFile(config, "/app/template/app/template/repository/repository.go", "app/template/repository/repository.go", cache.Repository, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Service, err = generateTemplateFile(config, "/app/template/app/template/service/service.go", "app/template/service/service.go", cache.Service, isUseCache)
		return
	})
	runTask(&wg, errChan, func() (err error) {
		cache.Template, err = generateTemplateFile(config, "/app/template/app/template/template.go", "app/template/template.go", cache.Template, isUseCache)
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
