package mod

import (
	"fmt"

	"github.com/ryanadiputraa/ggen/internal/util"
)

func Write(mod, name string) error {
	path := fmt.Sprintf("%v", name)
	return util.WriteToFile(template(mod), path)
}

func template(mod string) string {
	return fmt.Sprintf(`module %v

go 1.21.4`, mod)
}
