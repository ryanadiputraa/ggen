package module

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/ryanadiputraa/ggen/v2/pkg/logger"
)

func SetupProject(projectName, goMod string, log logger.Logger) (err error) {
	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os error: %v", err.Error())
	}

	if err = os.Chdir(projectName); err != nil {
		return fmt.Errorf("os error: %v", err.Error())
	}
	defer func() {
		if err := os.Chdir(originalDir); err != nil {
			log.Error("os error: %v", err.Error())
		}
	}()

	if err = cleanup(); err != nil {
		return fmt.Errorf("os error: %v", err.Error())
	}

	return initModule(goMod, log)
}

func cleanup() (err error) {
	if err = os.RemoveAll(".git"); err != nil {
		return
	}
	if err = os.RemoveAll(".github"); err != nil {
		return
	}
	if err = os.Remove("go.mod"); err != nil {
		return
	}
	return os.Remove("go.sum")
}

func initModule(goMod string, log logger.Logger) (err error) {
	var out bytes.Buffer

	gitCmd := exec.Command("git", "init")
	gitCmd.Stdout = &out
	gitCmd.Stderr = &out

	if err = gitCmd.Run(); err != nil {
		log.Error(out.String())
		return
	}
	log.Info(out.String())
	out.Reset()

	modInit := exec.Command("go", "mod", "init", goMod)
	modInit.Stdout = &out
	modInit.Stderr = &out

	if err = modInit.Run(); err != nil {
		log.Error(out.String())
		return
	}
	log.Info(out.String())
	out.Reset()

	modTidy := exec.Command("go", "mod", "tidy")
	modTidy.Stdout = &out
	modTidy.Stderr = &out

	if err = modTidy.Run(); err != nil {
		log.Error(out.String())
		return
	}
	log.Info(out.String())
	out.Reset()

	return
}
