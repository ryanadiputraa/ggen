package template

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	templatePath = "https://ggen.s3.ap-southeast-1.amazonaws.com/ggen-template"
	extension    = ".zip"
	tag          = "v1.0.0"
)

func FetchTemplate(projectName, goMod string) error {
	zipfile := projectName + ".zip"
	url := templatePath + "@" + tag + extension
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("network error: %v", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetch error, status: %v", resp.Status)
	}

	out, err := os.Create(zipfile)
	if err != nil {
		return fmt.Errorf("template error: %v", err.Error())
	}

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("template error: %v", err.Error())
	}

	if err := extractZipFile(zipfile, projectName, goMod); err != nil {
		return fmt.Errorf("template error: %v", err.Error())
	}

	if err := os.Remove(zipfile); err != nil {
		return fmt.Errorf("template error: %v", err.Error())
	}
	return nil
}

func extractZipFile(filename, dest, goMod string) (err error) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return
	}

	for _, f := range r.File {
		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			// Create directories as needed
			os.MkdirAll(path, os.ModePerm)
			continue
		}

		// Create file
		if err = os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return
		}
		outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		content, err := replaceTemplateModule(rc, dest, goMod)
		if err != nil {
			fmt.Println("err replace", err.Error())
			return err
		}
		if _, err := outFile.Write(content); err != nil {
			return err
		}
	}
	return
}

func replaceTemplateModule(rc io.ReadCloser, projectName, goMod string) (content []byte, err error) {
	templateName := "<ggen-template>"
	templateModule := "github.com/ryanadiputraa/ggen-template"

	var buf bytes.Buffer
	if _, err = io.Copy(&buf, rc); err != nil {
		return
	}

	modifiedContent := strings.ReplaceAll(buf.String(), templateName, projectName)
	modifiedContent = strings.ReplaceAll(modifiedContent, templateModule, goMod)
	content = []byte(modifiedContent)
	return
}
