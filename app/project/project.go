package project

import (
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ryanadiputraa/ggen/config"
)

const (
	userPermission = fs.FileMode(0700)
	templateURL    = "https://raw.githubusercontent.com/ryanadiputraa/ggen/refactor/generate-template"
)

func GenerateProjectTempalate(config *config.Config) (err error) {
	if err = writeConfigFile(config); err != nil {
		return
	}
	return
}

// tmplPath is the file path from templateURL base path
func generateTemplateFile(config *config.Config, tmplPath, destPath string) (err error) {
	content, err := fetchTemplate(tmplPath)
	if err != nil {
		return
	}
	return writeFile(config, content, destPath)
}

func writeFile(config *config.Config, content, destPath string) (err error) {
	modifiedName := strings.Replace(string(content), "%NAME%", config.ProjectName, -1)
	modifiedMod := strings.Replace(string(modifiedName), "%MOD%", config.GoMod, -1)

	err = os.WriteFile(filepath.Join(config.ProjectDir, destPath), []byte(modifiedMod), 0644)
	return
}

func fetchTemplate(path string) (content string, err error) {
	resp, err := http.Get(templateURL + path)
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
