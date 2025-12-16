package logger_config

import "github.com/NoOl01/velog/pkg/velog/velog_config"

type LoggerFunc struct {
	Timestamp bool
	Caller    bool
}

var LoggerFuncConfig *LoggerFunc
var ApiConfig *velog_config.Config
