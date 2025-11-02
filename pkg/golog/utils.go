package golog

import (
	"github.com/NoOl01/golog/pkg/golog/golog_config"
	loggerconfig "github.com/NoOl01/golog/pkg/golog/internal/config"
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

var defaultConfig = loggerConfig{
	Default: &logger.Logger{},
}

func Start(config *golog_config.Config) DefaultLogger {
	format.Format(config.Format, config.Literal)

	loggerconfig.LoggerFuncConfig = &loggerconfig.LoggerFunc{
		Timestamp: false,
		Caller:    false,
	}
	format.AutoDisable()

	if loggerconfig.LoggerFuncConfig.Timestamp {
		go logger.StartTickerTimestamp()
		<-logger.TimestampDone
	}

	logger.StartConsoleLog()

	return defaultConfig.Default
}

func Stop() {
	logger.StopConsoleLog()
}
