package module

import (
	"fmt"
	"os"
	"os/exec"
)

func NewModule(name, module string) (err error) {
	originalDir, err := os.Getwd()
	if err != nil {
		return
	}
	defer os.Chdir(originalDir)

	if err = createDirectory(name); err != nil {
		return
	}
	if err = os.Chdir(name); err != nil {
		return err
	}

	c := exec.Command("go", "mod", "init", module)
	err = c.Run()
	return
}

func createDirectory(dirName string) (err error) {
	_, err = os.Stat(dirName)
	if os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0700)
		return
	}

	if err != nil {
		return fmt.Errorf("fail to check directory status: %v", err)
	}
	return fmt.Errorf("directory \"%v\" already exists", dirName)
}
