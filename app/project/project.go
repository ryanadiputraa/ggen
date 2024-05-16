package project

import (
	"io"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/ryanadiputraa/ggen/config"
)

const (
	userPermission = fs.FileMode(0700)
	// TODO: change to main branch
	templateURL = "https://raw.githubusercontent.com/ryanadiputraa/ggen/refactor/generate-template"
)

func GenerateProjectTempalate(config *config.Config) (err error) {
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

// tmplPath is the file path from templateURL base path
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
