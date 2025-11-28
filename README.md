# GoLog

Lightweight and fast logger for Go with flexible formatting support

## Features

- Flexible configuration of the log output format
- High productivity and speed of work
- Support for multiple levels of logs (Info, Warn, Error, Debug, Fatal)
- Without third-party dependencies

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
    config := &golog_config.Config{
        Format: "${name} ${l} ${content} ${l} ${level} ${l} ${timestamp} ${l} ${caller}",
		Literal: " | ",
    }
    log := golog.Start(config)
    defer golog.Stop()
    
    log.Info("test", "hellow world")
    
    
}
```

## Benchmark

| Operation | Iterations | Speed | Memory | Allocations |
|:--:|:--:|:--:|:--:|:--:|
`BenchmarkLogToConsole | 7000000 | 150 ns/op | 0 B/op | 0 allocs/op

## Format placeholders

- ${level} - Log level
- ${name} - Log name
- ${content} - Log message
- ${timestamp} - Time of the log call
- ${caller} - The place where the log was called
- ${l} - Literal