package logger

import (
	"fmt"
	"runtime"
)

type Logger struct{}

const (
	INFO    string = "INFO"
	DEBUG   string = "DEBUG"
	WARNING string = "WARNING"
	ERROR   string = "ERROR"
	PANIC   string = "PANIC"
)

func (l *Logger) Info(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}

	Log(msg, INFO, file, line)
}

func (l *Logger) Debug(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}

	Log(msg, DEBUG, file, line)
}

func (l *Logger) Warn(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}

	Log(msg, WARNING, file, line)
}

func (l *Logger) Error(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}

	Log(msg, ERROR, file, line)
}

func (l *Logger) Panic(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}

	Log(msg, PANIC, file, line)
}
