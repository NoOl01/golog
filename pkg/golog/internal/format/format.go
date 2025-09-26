package format

import (
	"fmt"
	"github.com/NoOl01/golog/pkg/golog/golog_config"
	"github.com/NoOl01/golog/pkg/golog/golog_errs"
	"github.com/NoOl01/golog/pkg/golog/internal/tokens"
	"regexp"
	"strings"
)

var LevelToBytes = map[golog_config.LogLevel][]byte{
	golog_config.INFO:    []byte("INFO"),
	golog_config.WARNING: []byte("WARNING"),
	golog_config.ERROR:   []byte("ERROR"),
	golog_config.DEBUG:   []byte("DEBUG"),
	golog_config.PANIC:   []byte("PANIC"),
}

var L = map[string][]byte{}

type LogFormat []tokens.TokenType

var LogFormatTokens *LogFormat

func Format(format, literal string) {
	LogFormatTokens = &LogFormat{}

	L["l"] = []byte(literal)

	if err := ParseFormat(format, LogFormatTokens); err != nil {
		fmt.Println("ParseFormat(format) failed:", err)
	}
}

func ParseFormat(format string, logFormat *LogFormat) error {
	var tokenRegex = regexp.MustCompile(`\$\{([a-zA-Z0-9_]+)}`)
	formatTokens := tokenRegex.FindAllStringSubmatch(format, -1)

	for _, token := range formatTokens {
		token[1] = strings.TrimSpace(token[1])
		if token[1] == "" {
			continue
		}

		switch token[1] {
		case "name":
			*logFormat = append(*logFormat, tokens.TokenName)
		case "content":
			*logFormat = append(*logFormat, tokens.TokenContent)
		case "level":
			*logFormat = append(*logFormat, tokens.TokenLevel)
		case "timestamp":
			*logFormat = append(*logFormat, tokens.TokenTimestamp)
		case "caller":
			*logFormat = append(*logFormat, tokens.TokenCaller)
		case "l":
			*logFormat = append(*logFormat, tokens.TokenLiteral)
		default:
			return golog_errs.UnknownFormat
		}
	}

	return nil
}
