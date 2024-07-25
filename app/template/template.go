package template

import (
	"archive/zip"
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
	tag          = "v0.0.6"
)

func FetchTemplate(projectName string) error {
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

	if err := extractZipFile(zipfile, projectName); err != nil {
		return fmt.Errorf("template error: %v", err.Error())
	}

	if err := os.Remove(zipfile); err != nil {
		return fmt.Errorf("template error: %v", err.Error())
	}
	return nil
}

func extractZipFile(filename, dest string) error {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return err
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
		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err
		}
		outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		_, err = io.Copy(outFile, rc)

		// Close the file without deferring to avoid too many open files error
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}

	return nil
}
