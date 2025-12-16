package logger_config

import "github.com/NoOl01/velog/pkg/golog/golog_config"

type LoggerFunc struct {
	Timestamp bool
	Caller    bool
}

var LoggerFuncConfig *LoggerFunc
var ApiConfig *golog_config.Config
