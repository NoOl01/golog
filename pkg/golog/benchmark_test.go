package golog

import (
	"github.com/NoOl01/golog/pkg/golog/golog_config"
	"github.com/NoOl01/golog/pkg/golog/internal/format"
	"github.com/NoOl01/golog/pkg/golog/internal/logger"
	"testing"
)

func BenchmarkLogToConsole(b *testing.B) {
	golog_config.Format = "${name} ${l} ${content} ${l} ${level} ${l} ${timestamp} ${l} ${caller}"
	log := Start()

	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info("test", "hello world")
		}
	})

	Stop()
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
			logger.GetCaller()
		}
	})
}
