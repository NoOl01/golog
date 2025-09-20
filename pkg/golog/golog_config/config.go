package golog_config

type LogLevel int
type Color struct {
	R, G, B uint8
}
type Rotation struct {
	AutoDelete     bool
	AutoDeleteDays int
	Compress       bool
	CompressDays   int
}

type FileConfig struct {
	FileExtension string
	FilePath      string
	FileMaxSize   int
}

type LevelConfig struct {
	LogLevel LogLevel
	Color    Color
	Format   string
	Rotation Rotation
}

var (
	White  = Color{255, 255, 255}
	Red    = Color{255, 0, 0}
	Green  = Color{0, 255, 0}
	Blue   = Color{0, 0, 255}
	Yellow = Color{255, 255, 0}
	Orange = Color{255, 165, 0}
)

const (
	FTimestamp  = "timestamp"
	FColor      = "color"
	FName       = "name"
	FContent    = "content"
	FCallerFile = "caller_file"
	FCallerLine = "caller_line"
)

const (
	INFO LogLevel = iota
	DEBUG
	WARNING
	ERROR
	PANIC
)
