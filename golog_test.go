package main

import (
	"github.com/NoOl01/golog/pkg/golog"
	"testing"
)

func BenchmarkLog(b *testing.B) {
	logger := golog.Config.Default

	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("hello world")
		}
	})
}

func BenchmarkLogWithBuff(b *testing.B) {
	loggerBuff := golog.Config.Buff

	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			loggerBuff.InfoB("hello world")
		}
	})
}
