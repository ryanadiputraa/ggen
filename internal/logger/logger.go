package logger

import (
	"fmt"

	"github.com/ryanadiputraa/ggen/internal/writer"
)

func Write(mod, name string) error {
	path := fmt.Sprintf("%v/pkg/logger", name)
	if err := writer.CreateDirectory(path); err != nil {
		return err
	}
	return writer.WriteToFile(template(mod), path, "logger.go")
}

func template(mod string) string {
	return `package logger

import (
    "log"
    "os"
    "time"
)

type Logger interface {
    Info(v ...any)
    Warn(v ...any)
    Error(v ...any)
    Fatal(v ...any)
}

type logger struct {
    log *log.Logger
}

func NewLogger() Logger {
    location := time.UTC

    l := log.New(os.Stdout, "", log.Ldate|log.Ltime)
    l.SetFlags(l.Flags() | log.LUTC)
    l.SetPrefix("[" + location.String() + "] ")
    return &logger{log: l}
}

func (l *logger) Info(v ...any) {
    l.log.Println(v...)
}

func (l *logger) Infow(msg string, keyAndValues ...any) {
    l.log.Println(msg, keyAndValues)
}

func (l *logger) Warn(v ...any) {
    l.log.Println(v...)
}

func (l *logger) Error(v ...any) {
    l.log.Println(v...)
}

func (l *logger) Fatal(v ...any) {
    l.log.Fatal(v...)
}`
}
