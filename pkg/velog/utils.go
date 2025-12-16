package velog

import (
	"github.com/NoOl01/velog/internal/format"
	logger2 "github.com/NoOl01/velog/internal/logger"
	"github.com/NoOl01/velog/internal/logger/console_logger"
	"github.com/NoOl01/velog/internal/logger_config"
	"github.com/NoOl01/velog/pkg/velog/velog_config"
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
	Default: &logger2.Logger{},
}

func Start(config *velog_config.Config) DefaultLogger {
	logger_config.ApiConfig = config
	format.Format(config.Format, config.Literal)

	logger_config.LoggerFuncConfig = &logger_config.LoggerFunc{
		Timestamp: false,
		Caller:    false,
	}
	format.AutoDisable()

	if logger_config.LoggerFuncConfig.Timestamp {
		go logger2.StartTickerTimestamp()
		<-logger2.TimestampDone
	}

	console_logger.StartConsoleLog()

	return defaultConfig.Default
}

func Stop() {
	console_logger.StopConsoleLog()
}
