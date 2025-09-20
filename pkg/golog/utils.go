package golog

import (
	"fmt"
	"github.com/NoOl01/golog/pkg/golog/internal/format"
	"github.com/NoOl01/golog/pkg/golog/internal/logger"
)

type DefaultLogger interface {
	Info(msg string)
	Debug(msg string)
	Warn(msg string)
	Error(msg string)
	Panic(msg string)
}

type LogWithBuff interface {
	InfoB(msg string)
	DebugB(msg string)
	WarnB(msg string)
	ErrorB(msg string)
	PanicB(msg string)
}

type loggerConfig struct {
	Default DefaultLogger
	Buff    LogWithBuff
}

var config = loggerConfig{
	Default: &logger.Logger{},
	Buff:    &logger.Buff{},
}

func Start() DefaultLogger {

	return config.Default
}

func StartBuff() LogWithBuff {
	return config.Buff
}

func Test(text string) {
	var form format.LogFormat
	if err := format.Format(text, &form); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(form)
}
