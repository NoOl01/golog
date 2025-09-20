package logger

import (
	"bytes"
	"github.com/NoOl01/golog/pkg/golog/golog_config"
	"github.com/NoOl01/golog/pkg/golog/internal/buffer"
	"github.com/NoOl01/golog/pkg/golog/internal/format"
	"github.com/NoOl01/golog/pkg/golog/internal/tokens"
	"os"
	"time"
)

type LogData struct {
	Level     []byte
	Timestamp []byte
	Caller    []byte
	Literal   []byte
	Name      string
	Content   string
}

type Logger struct{}

func Log(name, msg string, level golog_config.LogLevel) {
	buf := buffer.GetBuffer()
	buf.Reset()
	defer buffer.PutBuffer(buf)

	data := &LogData{
		Name:      name,
		Content:   msg,
		Level:     format.LevelToBytes[level],
		Timestamp: []byte(time.Now().Format(time.RFC3339)),
		Caller:    []byte{},
		Literal:   format.L["l"],
	}

	for _, logFormat := range *format.LogFormatTokens {
		WriteByTokens(buf, logFormat, data)
	}

	buf.WriteByte('\n')

	os.Stdout.Write(buf.Bytes())
}

func (l *Logger) Info(name, msg string)  { Log(name, msg, golog_config.INFO) }
func (l *Logger) Debug(name, msg string) { Log(name, msg, golog_config.DEBUG) }
func (l *Logger) Warn(name, msg string) {
	Log(name, msg, golog_config.WARNING)
}
func (l *Logger) Error(name, msg string) { Log(name, msg, golog_config.ERROR) }
func (l *Logger) Panic(name, msg string) { Log(name, msg, golog_config.PANIC) }

func WriteByTokens(buf *bytes.Buffer, token tokens.TokenType, data *LogData) {
	switch token {
	case tokens.TokenName:
		buf.WriteString(data.Name)
	case tokens.TokenContent:
		buf.WriteString(data.Content)
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
