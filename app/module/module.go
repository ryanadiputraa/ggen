package module

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ryanadiputraa/ggen/v2/config"
)

func NewModule(config *config.Config) (err error) {
	if err = createDirectory(config.ProjectName); err != nil {
		return
	}
	if err = os.Chdir(config.ProjectName); err != nil {
		return
	}

	// git init
	c := exec.Command("git", "init")
	if err = c.Run(); err != nil {
		return
	}

	// init go module
	c = exec.Command("go", "mod", "init", config.GoMod)
	err = c.Run()
	return
}

func TidyGoMod() (err error) {
	c := exec.Command("go", "mod", "tidy")
	return c.Run()
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
