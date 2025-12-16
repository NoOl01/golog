package logger

import (
	"runtime"
	"strconv"
	"sync"

	"github.com/nool01/velog/internal/logger_config"
)

var (
	callerCache sync.Map
)

// GetCaller todo Fix caller
func GetCaller() []byte {
	if !logger_config.LoggerFuncConfig.Caller {
		return nil
	}

	var pcs [1]uintptr
	runtime.Callers(4, pcs[:])

	if v, ok := callerCache.Load(pcs[0]); ok {
		return v.([]byte)
	}

	file, line := runtime.FuncForPC(pcs[0]).FileLine(pcs[0])

	b := make([]byte, 0, 64)
	b = append(b, file...)
	b = append(b, ':')
	b = strconv.AppendInt(b, int64(line), 10)

	callerCache.Store(pcs[0], b)
	return b
}
