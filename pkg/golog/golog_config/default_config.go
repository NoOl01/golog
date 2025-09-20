package golog_config

import "fmt"

var (
	Format  = fmt.Sprintf("${%s} ${l} ${%s}} ${l} ${%s}", FName, FContent, FTimestamp)
	Literal = " | "
	InfoCfg = LevelConfig{
		LogLevel: INFO,
		Color:    White,
		Rotation: Rotation{
			AutoDelete:     true,
			AutoDeleteDays: 30,
			Compress:       true,
			CompressDays:   5,
		},
	}
	DebugCfg = LevelConfig{
		LogLevel: DEBUG,
		Color:    Green,
		Rotation: Rotation{
			AutoDelete:     true,
			AutoDeleteDays: 30,
			Compress:       true,
			CompressDays:   5,
		},
	}
	WarningCfg = LevelConfig{
		LogLevel: WARNING,
		Color:    Yellow,
		Rotation: Rotation{
			AutoDelete:     true,
			AutoDeleteDays: 30,
			Compress:       true,
			CompressDays:   5,
		},
	}
	ErrorCfg = LevelConfig{
		LogLevel: ERROR,
		Color:    Orange,
		Rotation: Rotation{
			AutoDelete:     true,
			AutoDeleteDays: 30,
			Compress:       true,
			CompressDays:   5,
		},
	}
	PanicCfg = LevelConfig{
		LogLevel: PANIC,
		Color:    Red,
		Rotation: Rotation{
			AutoDelete:     true,
			AutoDeleteDays: 30,
			Compress:       true,
			CompressDays:   5,
		},
	}
)

var DefaultFileConfig = FileConfig{
	FileExtension: "txt",
	FilePath:      "./",
	FileMaxSize:   10,
}
