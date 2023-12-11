package domain

import (
	"fmt"

	"github.com/ryanadiputraa/ggen/internal/writer"
)

func Write(mod, name string) error {
	path := fmt.Sprintf("%v/internal/domain/test", name)
	if err := writer.CreateDirectory(path); err != nil {
		return err
	}

	return writer.WriteToFile(template(mod), path, "test.go")
}

func template(mod string) string {
	return `package test

type Test struct {
    Message string ` + "`json:\"message\"`" + `
}`
}
