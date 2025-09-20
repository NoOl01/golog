package golog

import (
	"github.com/NoOl01/golog/pkg/golog/golog_config"
	"github.com/NoOl01/golog/pkg/golog/internal/format"
	"github.com/NoOl01/golog/pkg/golog/internal/logger"
)

type DefaultLogger interface {
	Info(name, msg string)
	Debug(name, msg string)
	Warn(name, msg string)
	Error(name, msg string)
	Panic(name, msg string)
}

type loggerConfig struct {
	Default DefaultLogger
}

var config = loggerConfig{
	Default: &logger.Logger{},
}

func Start() DefaultLogger {
	format.Format(golog_config.Format, golog_config.Literal)

	return config.Default
}
