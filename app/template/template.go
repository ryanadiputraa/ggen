package template

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	filepath  = "https://ggen.s3.ap-southeast-1.amazonaws.com/ggen-template"
	extension = ".zip"
	tag       = "v0.0.6"
)

func FetchTemplate(projectName string) error {
	url := filepath + "@" + tag + extension
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("network error: %v", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetch error, status: %v", resp.Status)
	}

	out, err := os.Create(projectName + ".zip")
	if err != nil {
		return fmt.Errorf("template error: %v", err.Error())
	}

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("template error: %v", err.Error())
	}

	return nil
}
