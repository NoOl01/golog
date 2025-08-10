package json

type Json struct {
	LogLevels      []LogLevel     `json:"log_levels"`
	FileConfig     FileConfig     `json:"file_config"`
	DatabaseConfig DatabaseConfig `json:"database_config"`
}

type LogLevel struct {
	Name     string   `json:"name"`
	Debug    bool     `json:"debug"`
	Format   string   `json:"format"`
	Color    string   `json:"color"`
	Outputs  []string `json:"outputs"`
	Rotation Rotation `json:"rotation"`
}

type Rotation struct {
	AutoDelete     bool `json:"auto_delete"`
	AutoDeleteDays int  `json:"auto_delete_days"`
	Compress       bool `json:"compress"`
	CompressDays   int  `json:"compress_days"`
}

type FileConfig struct {
	FileExtension string `json:"file_extension"`
	FilePath      string `json:"file_path"`
	FileMaxSize   int    `json:"file_max_size"`
}

type DatabaseConfig struct {
	Driver     string     `json:"driver"`
	Dsn        string     `json:"dsn"`
	Table      string     `json:"table"`
	TableModel TableModel `json:"table_model"`
}

type TableModel map[string]map[string]interface{}
