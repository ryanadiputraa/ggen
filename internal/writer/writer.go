package writer

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateDirectory(dir string) error {
	return os.MkdirAll(dir, 0755)
}

func WriteToFile(content, filePath, fileName string) error {
	c := exec.Command("bash", "-c", fmt.Sprintf("echo '%s' > %s/%s", content, filePath, fileName))
	return c.Run()
}
