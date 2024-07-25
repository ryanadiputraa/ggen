package logger

import (
	"log"
	"os"
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
	l := log.New(os.Stdout, "", 0)
	return &logger{log: l}
}

func (l *logger) Info(v ...any) {
	l.log.Println(v...)
}

func (l *logger) Warn(v ...any) {
	l.log.Println(v...)
}

func (l *logger) Error(v ...any) {
	l.log.Println(v...)
}

func (l *logger) Fatal(v ...any) {
	l.log.Fatal(v...)
}
