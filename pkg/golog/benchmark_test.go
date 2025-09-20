package golog

import (
	"github.com/NoOl01/golog/pkg/golog/internal/format"
	"testing"
)

func BenchmarkLogToConsole(b *testing.B) {
	loggerBuff := Start()

	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			loggerBuff.Info("test", "hello world")
		}
	})
}

func BenchmarkFormat(b *testing.B) {
	b.ReportAllocs()

	var logFormat format.LogFormat

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			format.ParseFormat("${name}, ${content}", &logFormat)
		}
	})
}

func BenchmarkCaller(b *testing.B) {
	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			format.Caller()
		}
	})
}
