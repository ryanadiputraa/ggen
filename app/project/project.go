package project

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ryanadiputraa/ggen/app/cache"
	"github.com/ryanadiputraa/ggen/config"
	"github.com/ryanadiputraa/ggen/pkg/github"
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

	tag, err := checkTag()
	if err != nil {
		log.Println("fail to get latest tag: ", err)
	} else {
		isUseCache = c.Tag == tag
	}

	if !isUseCache {
		fmt.Println("fetching template... (skipping cache")
		cache.StoreCache(cache.Cache{Tag: tag})
	}

	if err = writeConfigFile(config); err != nil {
		return
	}
	if err = writeCMD(config); err != nil {
		return
	}
	if err = writeServer(config); err != nil {
		return
	}
	if err = writeApp(config); err != nil {
		return
	}
	return writePkg(config)
}

// tmplPath is the file path from github.TemplateURL base path
func generateTemplateFile(config *config.Config, tmplPath, destPath string) (err error) {
	content, err := fetchTemplate(tmplPath)
	if err != nil {
		return
	}
	return writeFile(config, content, destPath)
}

func writeFile(config *config.Config, content, destPath string) (err error) {
	modifiedMod := strings.Replace(string(content), "github.com/ryanadiputraa/ggen/app/template", config.GoMod, -1)

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
