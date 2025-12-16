package pkg

import (
	"testing"

	"github.com/NoOl01/velog/pkg/velog"
	"github.com/NoOl01/velog/pkg/velog/velog_config"
)

func BenchmarkLogToConsole(b *testing.B) {
	config := &velog_config.Config{
		Format:  "${name} ${l} ${content} ${l} ${level} ${l} ${timestamp} ${l} ${caller}",
		Literal: " || ",
	}
	log := velog.Start(config)

	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info("test", "hello world")
		}
	})

	velog.Stop()
}
