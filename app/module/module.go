package module

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ryanadiputraa/ggen/config"
)

func NewModule(config *config.Config) (err error) {
	ggenDir, err := os.Getwd()
	if err != nil {
		return
	}
	defer os.Chdir(ggenDir)

	if err = createDirectory(config.ProjectName); err != nil {
		return
	}
	if err = os.Chdir(config.ProjectName); err != nil {
		return err
	}

	// set directory config
	projectDir, err := os.Getwd()
	if err != nil {
		return
	}
	config.GgenDir = ggenDir
	config.ProjectDir = projectDir

	// init go module
	c := exec.Command("go", "mod", "init", config.GoMod)
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
