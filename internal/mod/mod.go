package mod

import (
	"fmt"

	"github.com/ryanadiputraa/ggen/internal/writer"
)

func Write(mod, name string) error {
	return writer.WriteToFile(template(mod), name, "go.mod")
}

func template(mod string) string {
	return fmt.Sprintf(`module %v

go 1.21.4`, mod)
}
