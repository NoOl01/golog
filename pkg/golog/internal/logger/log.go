package logger

import (
	"fmt"
	"time"
)

func Log(msg string, level string, file string, line int) {
	timestamp := time.Now().Format("2006-01-2 15:04:05")
	switch level {
	case INFO:
		fmt.Printf("%s | \033[38;2;255;255;255m%-7s\033[0m | %-7s | %s %d\n", timestamp, level, msg, file, line)
	case DEBUG:
		fmt.Printf("%s | \033[38;2;0;255;0m%-7s\033[0m | %-7s | %s %d\n", timestamp, level, msg, file, line)
	case WARNING:
		fmt.Printf("%s | \033[38;2;255;255;0m%-7s\033[0m | %-7s | %s %d\n", timestamp, level, msg, file, line)
	case ERROR:
		fmt.Printf("%s | \033[38;2;255;165;0m%-7s\033[0m | %-7s | %s %d\n", timestamp, level, msg, file, line)
	case PANIC:
		fmt.Printf("%s | \033[38;2;255;0;0m%-7s\033[0m | %-7s | %s %d\n", timestamp, level, msg, file, line)
	}
}
