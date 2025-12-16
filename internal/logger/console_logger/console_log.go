package console_logger

import (
	"bufio"
	"bytes"
	"os"
	"sync"
	"time"

	"github.com/nool01/velog/internal/buffer"
)

var (
	logWg            sync.WaitGroup
	LogBufferChannel = make(chan *bytes.Buffer, 100)
	logQuit          = make(chan struct{})
	once             sync.Once
)

func StartConsoleLog() {
	logWg.Add(1)
	go consoleLogTicker()
}

func consoleLogTicker() {
	defer logWg.Done()

	writer := bufio.NewWriter(os.Stdout)
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	const maxBatchSize = 1024
	count := 0

	for {
		select {
		case buf := <-LogBufferChannel:
			n, _ := writer.Write(buf.Bytes())
			count += n
			buffer.PutBuffer(buf)

			if count >= maxBatchSize {
				writer.Flush()
				count = 0
			}
		case <-ticker.C:
			if count > 0 {
				writer.Flush()
				count = 0
			}
		case <-logQuit:
			for {
				select {
				case buf := <-LogBufferChannel:
					writer.Write(buf.Bytes())
					buffer.PutBuffer(buf)
				default:
					writer.Flush()
					return
				}
			}
		}
	}
}

func StopConsoleLog() {
	once.Do(func() {
		close(logQuit)
	})
	logWg.Wait()
}
