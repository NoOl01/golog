package format

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func formatFromHex(hex string) (string, error) {
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
