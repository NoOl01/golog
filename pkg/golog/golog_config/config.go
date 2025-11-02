package golog_config

type LogLevel int

type Config struct {
	Format  string
	Literal string
}

const (
	INFO LogLevel = iota
	DEBUG
	WARNING
	ERROR
	PANIC
)
