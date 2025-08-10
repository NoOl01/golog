package golog

import "github.com/NoOl01/golog/pkg/golog/internal/logger"

type DefaultLogger interface {
	Info(msg string)
	Debug(msg string)
	Warn(msg string)
	Error(msg string)
	Panic(msg string)
}

type config struct {
	Default DefaultLogger
}

var Config = config{
	Default: &logger.Logger{},
}
