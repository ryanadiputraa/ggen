package cache

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type Cache struct {
	Tag string `json:"tag"`
	// raw code template cache
	Config             string `json:"config"`
	Env                string `json:"env"`
	Compose            string `json:"compose"`
	Main               string `json:"main"`
	CMD                string `json:"cmd"`
	Server             string `json:"server"`
	Handler            string `json:"handler"`
	Healthcheck        string `json:"healthcheck"`
	HealthcheckHandler string `json:"healthcheck_hadler"`
	Cors               string `json:"cors"`
	Throttle           string `json:"throttle"`
	Timeout            string `json:"timeout"`
	Postgres           string `json:"postgres"`
	Respwr             string `json:"respwr"`
	Logger             string `json:"logger"`
	Validator          string `json:"validator"`
	Gitignore          string `json:"gitignore"`
	Air                string `json:"air"`
	Makefile           string `json:"makefile"`
	Readme             string `json:"readme"`
	SetupTest          string `json:"setup_test"`
	HealthcheckTest    string `json:"healthcheck_test"`
}

const (
	userPermission = fs.FileMode(0700)
	cacheFile      = "cache.json"
	InitTag        = "v0.0.0"
)

func GetCache() (cache *Cache, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	cachePath := filepath.Join(homeDir, ".ggen")
	_, err = os.Stat(cachePath)
	if os.IsNotExist(err) {
		if err = os.Mkdir(cachePath, userPermission); err != nil {
			return
		}
		// return init tag to continue proccess
		return &Cache{Tag: InitTag}, nil
	}
	if err != nil {
		return &Cache{}, fmt.Errorf("fail to check cache status: %v", err)
	}

	content, err := os.ReadFile(filepath.Join(cachePath, cacheFile))
	if err != nil {
		if os.IsNotExist(err) {
			// return init tag to continue proccess
			return &Cache{Tag: InitTag}, nil
		}
		return
	}

	err = json.Unmarshal(content, &cache)
	return
}

func StoreCache(cache Cache) (err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return
	}

	cachePath := filepath.Join(homeDir, ".ggen")
	_, err = os.Stat(cachePath)
	if os.IsNotExist(err) {
		err = os.Mkdir(cachePath, userPermission)
		return
	}
	if err != nil {
		return fmt.Errorf("fail to check cache status: %v", err)
	}

	data, err := json.Marshal(cache)
	if err != nil {
		return
	}

	err = os.WriteFile(filepath.Join(cachePath, cacheFile), data, 0644)
	return
}
