package golog

import (
	"github.com/NoOl01/golog/pkg/golog/internal/format"
	"testing"
)

func BenchmarkLog(b *testing.B) {
	logger := Start()

	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("hello world")
		}
	})
}

func BenchmarkLogWithBuff(b *testing.B) {
	loggerBuff := StartBuff()

	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			loggerBuff.InfoB("hello world")
		}
	})
}

func BenchmarkFormat(b *testing.B) {
	b.ReportAllocs()

	var logFormat format.LogFormat

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			format.Format("${name}, ${content}", &logFormat)
		}
	})
}
