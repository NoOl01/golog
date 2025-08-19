package logger

import (
	"fmt"
	"github.com/NoOl01/golog/pkg/golog/golog_config"
	"runtime"
	"time"
)

type Logger struct{}

func Log(msg string, level golog_config.LogLevel) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}

	timestamp := time.Now().Format("2006-01-2 15:04:05")

	fmt.Printf("%s | \033[38;2;255;255;255m%-7d\033[0m | %-7s | %s %d\n", timestamp, level, msg, file, line)
}

func (l *Logger) Info(msg string)  { Log(msg, golog_config.INFO) }
func (l *Logger) Debug(msg string) { Log(msg, golog_config.DEBUG) }
func (l *Logger) Warn(msg string)  { Log(msg, golog_config.WARNING) }
func (l *Logger) Error(msg string) { Log(msg, golog_config.ERROR) }
func (l *Logger) Panic(msg string) { Log(msg, golog_config.PANIC) }
