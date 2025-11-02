package logger

import (
	"github.com/NoOl01/golog/pkg/golog/golog_config"
	"github.com/NoOl01/golog/pkg/golog/internal/buffer"
	"github.com/NoOl01/golog/pkg/golog/internal/format"
	"github.com/NoOl01/golog/pkg/golog/internal/logger/log_data"
	"github.com/NoOl01/golog/pkg/golog/internal/tokens"
	"github.com/NoOl01/golog/pkg/golog/internal/unsafe_conv"
)

type Logger struct{}

func (l *Logger) Log(name, msg string, level golog_config.LogLevel) {
	buf := buffer.GetBuffer()

	data := log_data.LogData{
		Name:      unsafe_conv.StringToBytes(name),
		Content:   unsafe_conv.StringToBytes(msg),
		Level:     format.LevelToBytes[level],
		Timestamp: timestamp,
		Caller:    GetCaller(),
		Literal:   format.L["l"],
	}

	for _, token := range *format.LogFormatTokens {
		switch token {
		case tokens.TokenName:
			buf.Write(data.Name)
		case tokens.TokenContent:
			buf.Write(data.Content)
		case tokens.TokenLevel:
			buf.Write(data.Level)
		case tokens.TokenTimestamp:
			buf.Write(data.Timestamp)
		case tokens.TokenCaller:
			buf.Write(data.Caller)
		case tokens.TokenLiteral:
			buf.Write(data.Literal)
		}
	}

	buf.WriteByte('\n')

	select {
	case logBufferChannel <- buf:
	default:
		buffer.PutBuffer(buf)
	}
}

func (l *Logger) Info(name, msg string)  { l.Log(name, msg, golog_config.INFO) }
func (l *Logger) Debug(name, msg string) { l.Log(name, msg, golog_config.DEBUG) }
func (l *Logger) Warn(name, msg string)  { l.Log(name, msg, golog_config.WARNING) }
func (l *Logger) Error(name, msg string) { l.Log(name, msg, golog_config.ERROR) }
func (l *Logger) Panic(name, msg string) { l.Log(name, msg, golog_config.PANIC) }
