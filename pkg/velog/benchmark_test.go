package velog

import (
	"testing"

	"github.com/NoOl01/velog/internal/format"
	"github.com/NoOl01/velog/internal/logger"
	"github.com/NoOl01/velog/pkg/velog/velog_config"
)

func BenchmarkLogToConsole(b *testing.B) {
	config := &velog_config.Config{
		Format:  "${name} ${l} ${content} ${l} ${level} ${l} ${timestamp} ${l} ${caller}",
		Literal: " || ",
	}
	log := Start(config)

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
