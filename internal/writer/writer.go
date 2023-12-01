package writer

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateDirectory(dir string) error {
	return os.MkdirAll(dir, 0755)
}

func WriteToFile(content, filePath, filename string) error {
	c := exec.Command("bash", "-c", fmt.Sprintf("echo '%s' > %s/%s", content, filePath, filename))
	return c.Run()
}
