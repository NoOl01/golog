package logger

import (
	"bytes"
	"github.com/NoOl01/golog/pkg/golog/internal/buffer"
	"runtime"
	"strconv"
	"sync"
)

var (
	fileBuffer = make([]byte, 0, 128)
	lineBuffer = make([]byte, 0, 10)
)

var callerPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func GetCaller() []byte {
	callerBuf := callerPool.Get().(*bytes.Buffer)
	callerBuf.Reset()

	pc := make([]uintptr, 15)
	runtime.Callers(3, pc)

	file, line := runtime.FuncForPC(pc[0]).FileLine(pc[0])
	callerBuf.WriteString(file)
	callerBuf.WriteByte(':')
	callerBuf.WriteString(strconv.Itoa(line))

	buffer.PutBuffer(callerBuf)
	return callerBuf.Bytes()
}
