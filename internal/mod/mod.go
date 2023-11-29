package mod

import (
	"fmt"
	"os/exec"
)

func Write(mod, name string) error {
	path := fmt.Sprintf("%v", name)
	write := exec.Command("bash", "-c", fmt.Sprintf("echo '%s' > %s/go.mod", template(mod), path))
	return write.Run()
}

func template(mod string) string {
	return fmt.Sprintf(`module %v

go 1.21.4`, mod)
}
