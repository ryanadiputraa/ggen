package project

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ryanadiputraa/ggen/v2/app/cache"
	"github.com/ryanadiputraa/ggen/v2/config"
	"github.com/ryanadiputraa/ggen/v2/pkg/github"
)

const (
	userPermission = fs.FileMode(0700)
)

func GenerateProjectTempalate(config *config.Config) (err error) {
	isUseCache := false

	fmt.Println("checking cache...")
	c, err := cache.GetCache()
	if err != nil {
		log.Println("fail to check cache: ", err)
		c.Tag = cache.InitTag
	}

	fmt.Println("checking latest tag...")
	tag, err := checkTag()
	if err != nil {
		log.Println("fail to get latest tag: ", err)
	} else {
		isUseCache = c.Tag == tag
	}

	wg := sync.WaitGroup{}
	templateErr := make(chan error, 5)

	// TOOD: refactor template

	runTask(&wg, templateErr, func() error { return writeApp(config, isUseCache, c) })
	runTask(&wg, templateErr, func() error { return writeMain(config, isUseCache, c) })
	runTask(&wg, templateErr, func() error { return writeConfig(config, isUseCache, c) })
	runTask(&wg, templateErr, func() error { return writePkg(config, isUseCache, c) })
	runTask(&wg, templateErr, func() error { return writeTest(config, isUseCache, c) })

	wg.Wait()
	close(templateErr)

	for e := range templateErr {
		if e != nil {
			err = e
		}
	}
	if err != nil {
		clenaup(config)
		return
	}

	if !isUseCache {
		c.Tag = tag
		cache.StoreCache(*c)
	}
	return
}

// tmplPath is the file path from github.TemplateURL base path, cache can be an empty string is isUseCache is false
func generateTemplateFile(config *config.Config, tmplPath, destPath, cache string, isUseCache bool) (generated string, err error) {
	var content string

	if !isUseCache {
		content, err = fetchTemplate(tmplPath)
		if err != nil {
			return
		}
	} else {
		content = cache
	}
	return content, writeFile(config, content, destPath)
}

func writeFile(config *config.Config, content, destPath string) (err error) {
	modifiedMod := strings.Replace(string(content), "github.com/ryanadiputraa/ggen/v2/app/template", config.GoMod, -1)
	err = os.WriteFile(destPath, []byte(modifiedMod), 0644)
	return
}

func fetchTemplate(path string) (content string, err error) {
	resp, err := http.Get(github.TemplateURL + path)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	content = string(body)
	return
}

func checkTag() (tag string, err error) {
	resp, err := http.Get(github.TagURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var tags []github.Tag
	if err = json.NewDecoder(resp.Body).Decode(&tags); err != nil {
		return
	}
	if len(tags) < 1 {
		return cache.InitTag, nil
	}

	return tags[0].Name, nil
}

func clenaup(config *config.Config) {
	_, err := os.Stat(config.OriginPath)
	if err != nil {
		// skip cleanup if project folder doesn't exists
		if os.IsNotExist(err) {
			return
		}
		return
	}

	if err := os.RemoveAll(filepath.Join(config.OriginPath, config.ProjectName)); err != nil {
		log.Println("fail to cleanup: ", err)
	}
}

func runTask(wg *sync.WaitGroup, ch chan<- error, task func() error) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := task(); err != nil {
			ch <- err
		}
	}()
}
