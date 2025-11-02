package logger

import (
	"sync"
	"time"
)

var (
	timestamp     = make([]byte, 0, 30)
	TimestampDone = make(chan struct{})
	onceTimestamp sync.Once
)

func StartTickerTimestamp() {
	updateTimestamp()
	onceTimestamp.Do(func() {
		close(TimestampDone)
	})
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		updateTimestamp()
	}
}

func updateTimestamp() {
	timestamp = time.Now().AppendFormat(timestamp[:0], time.RFC3339)
}
