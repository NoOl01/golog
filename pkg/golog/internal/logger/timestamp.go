package logger

import (
	"time"
)

var (
	timestamp = make([]byte, 0, 30)
)

func StartTickerTimestamp() {
	updateTimestamp()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		updateTimestamp()
	}
}

func updateTimestamp() {
	timestamp = time.Now().AppendFormat(timestamp[:0], time.RFC3339)
}
