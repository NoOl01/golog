package format

import (
	"errors"
	"fmt"
	"github.com/NoOl01/golog/pkg/golog/internal/tokens"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type LogFormat []tokens.Token

func Caller() (string, int) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("runtime.Caller() failed")
		return "", 0
	}

	return file, line
}

func Format() *LogFormat {
	return &LogFormat{}
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
			*logFormat = append(*logFormat, tokens.Token{Type: tokens.TokenName})
		case "content":
			*logFormat = append(*logFormat, tokens.Token{Type: tokens.TokenContent})
		case "level":
			*logFormat = append(*logFormat, tokens.Token{Type: tokens.TokenLevel})
		case "color":
			*logFormat = append(*logFormat, tokens.Token{Type: tokens.TokenColor})
		case "timestamp":
			*logFormat = append(*logFormat, tokens.Token{Type: tokens.TokenTimestamp})
		case "caller":
			*logFormat = append(*logFormat, tokens.Token{Type: tokens.TokenCaller})
		default:
			*logFormat = append(*logFormat, tokens.Token{Type: tokens.TokenLiteral})
		}
	}

	return nil
}

func HexToFormat(hex string) (string, error) {
	hex = strings.TrimPrefix(hex, "#")

	if len(hex) != 6 {
		return "", errors.New("invalid format")
	}

	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return "", err
	}
	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return "", err
	}
	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("\\033[38;2;%d;%d;%d", r, g, b), nil
}
