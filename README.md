# GoLog

Lightweight and fast logger for Go with flexible formatting support

## Installation

```
go get github.com/NoOl01/golog
```

## Quick start

### Simple example

```
package main

import "github.com/NoOl01/golog/pkg/golog"

func main() {
    log := golog.Start()
    
    log.Info("test", "hellow world")
    
    golog.Stop()
}
```

### Server example

```
func main() {
	golog_config.Literal = " || "
	golog_config.Format = "${level} ${l} ${name} ${l} ${content} ${l} ${timestamp}"

	log := golog.Start()
	defer golog.Stop()

	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		log.Info("http [ping]", "got request /ping")
		fmt.Fprintf(writer, "pong")
	})

	log.Info("server", "started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error("server", fmt.Sprintf("failed to startL %v", err))
	}
}
```

## Benchmark

### Benchmark test
```
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
```

### Benchmark result

`BenchmarkLogToConsole-12   6059786    169.0 ns/op  0 B/op   0 allocs/op`