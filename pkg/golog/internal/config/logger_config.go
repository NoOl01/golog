package config

type LoggerFunc struct {
	Timestamp bool
	Caller    bool
}

var LoggerFuncConfig *LoggerFunc
