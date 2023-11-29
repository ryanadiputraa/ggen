package util

import (
	"fmt"
	"os"
	"os/exec"
)

func MakeDirectory(dir string) error {
	return os.MkdirAll(dir, 0755)
}

func WriteToFile(content, filePath string) error {
	c := exec.Command("bash", "-c", fmt.Sprintf("echo '%s' > %s/main.go", content, filePath))
	return c.Run()
}
