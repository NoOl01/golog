package golog

import "github.com/NoOl01/golog/pkg/golog/internal/logger"

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

type config struct {
	Default DefaultLogger
	Buff    LogWithBuff
}

var Config = config{
	Default: &logger.Logger{},
	Buff:    &logger.Buff{},
}
