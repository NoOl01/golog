package logger

import (
	"sync"
	"testing"
	"time"
)

func TestUpdateTimestamp(t *testing.T) {
	timestamp = timestamp[:0]
	TimestampDone = make(chan struct{})
	onceTimestamp = sync.Once{}

	go StartTickerTimestamp()

	select {
	case <-TimestampDone:
		if len(timestamp) == 0 {
			t.Fatal("timestamp not updated")
		}
	case <-time.After(time.Second):
		t.Fatal("timestamp update timed out")
	}
}
