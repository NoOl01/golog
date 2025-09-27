package logger

import (
	"runtime"
	"strconv"
	"sync"
)

var (
	pcBuf      = make([]uintptr, 1)
	callerPool = sync.Pool{
		New: func() interface{} {
			buf := make([]byte, 0, 128)
			return &buf
		},
	}
)

func GetCaller() []byte {
	bufPtr := callerPool.Get().(*[]byte)
	buf := (*bufPtr)[:0]

	runtime.Callers(3, pcBuf)

	file, line := runtime.FuncForPC(pcBuf[0]).FileLine(pcBuf[0])

	buf = append(buf, file...)
	buf = append(buf, ':')
	buf = strconv.AppendInt(buf, int64(line), 10)

	*bufPtr = buf
	callerPool.Put(bufPtr)

	return buf
}
