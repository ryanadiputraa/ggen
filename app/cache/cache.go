package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type Cache struct {
	Tag string `json:"tag"`
}

const (
	userPermission = fs.FileMode(0700)
	cacheFile      = "cache.json"
	initTag        = "v0.0.0"
)

func GetCache() (cache Cache, err error) {
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
		return Cache{Tag: initTag}, nil
	}
	if err != nil {
		return Cache{}, fmt.Errorf("fail to check cache status: %v", err)
	}

	content, err := os.ReadFile(filepath.Join(cachePath, cacheFile))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// return init tag to continue proccess
			return Cache{Tag: initTag}, nil
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
