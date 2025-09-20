package logger

import (
	"fmt"
	"github.com/NoOl01/golog/pkg/golog/golog_config"
	"github.com/NoOl01/golog/pkg/golog/internal/buffer"
	"os"
	"runtime"
	"time"
)

type LogData struct {
	Name []byte
	Content []byte
	Level []byte
	Timestamp []byte
	Caller []byte
}

type Logger struct{}
type Buff struct{}

func Log(msg string, level golog_config.LogLevel) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return
	}

	timestamp := time.Now().Format("2006-01-2 15:04:05")

	fmt.Printf("%s | \033[38;2;255;255;255m%-7d\033[0m | %-7s | %s %d\n", timestamp, level, msg, file, line)
}

func LogWithBuffer(msg string, level golog_config.LogLevel) {
	buf := buffer.GetBuffer()
	defer buffer.PutBuffer(buf)

	//timestamp := time.Now().Format("2006-01-2 15:04:05")
	//_, err := fmt.Fprintf(buf, "%s | \033[38;2;255;255;255m%-7d\033[0m | %-7s | %s %d\n", timestamp, level, msg, file, line)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	buf.WriteString(msg)
	buf.WriteString("\n")

	_, _ = os.Stdout.Write(buf.Bytes())
}

func (l *Logger) Info(msg string)  { Log(msg, golog_config.INFO) }
func (l *Logger) Debug(msg string) { Log(msg, golog_config.DEBUG) }
func (l *Logger) Warn(msg string)  { Log(msg, golog_config.WARNING) }
func (l *Logger) Error(msg string) { Log(msg, golog_config.ERROR) }
func (l *Logger) Panic(msg string) { Log(msg, golog_config.PANIC) }

func (l *Buff) InfoB(msg string)  { LogWithBuffer(msg, golog_config.INFO) }
func (l *Buff) DebugB(msg string) { LogWithBuffer(msg, golog_config.DEBUG) }
func (l *Buff) WarnB(msg string)  { LogWithBuffer(msg, golog_config.WARNING) }
func (l *Buff) ErrorB(msg string) { LogWithBuffer(msg, golog_config.ERROR) }
func (l *Buff) PanicB(msg string) { LogWithBuffer(msg, golog_config.PANIC) }
