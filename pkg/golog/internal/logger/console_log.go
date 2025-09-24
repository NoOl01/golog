package logger

import (
	"bufio"
	"bytes"
	"github.com/NoOl01/golog/pkg/golog/internal/buffer"
	"os"
	"sync"
	"time"
)

var (
	logWg            sync.WaitGroup
	logBufferChannel = make(chan *bytes.Buffer, 100)
	logQuit          chan struct{}
)

var quit = make(chan struct{})

func StartConsoleLog() {

	logWg.Add(1)
	go consoleLogTicker()

	logWg.Wait()
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
		case buf := <-logBufferChannel:
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
		case <-quit:
			for {
				select {
				case buf := <-logBufferChannel:
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
	close(quit)
	logWg.Wait()
}
