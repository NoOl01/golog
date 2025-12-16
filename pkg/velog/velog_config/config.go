package velog_config

type LogLevel int
type Format int

const (
	Disabled Format = iota
	Text
	Json
)

type Config struct {
	Format  string
	Literal string
	Debug   bool
	Console OutputFormat
	File    FileOutput
}

type FileOutput struct {
	FileExtension string
	AutoDelete    int
	AutoArchive   int
	Output        OutputFormat
}

type OutputFormat struct {
	Format Format
	Json   string
}

const (
	INFO LogLevel = iota
	DEBUG
	WARNING
	ERROR
	PANIC
)
