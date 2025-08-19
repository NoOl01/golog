package golog_config

import "fmt"

var (
	InfoCfg = LevelConfig{
		LogLevel: INFO,
		Color:    White,
		Format:   fmt.Sprintf("%s | %s{%s} | %s | %s:%s", FTimestamp, FColor, FName, FContent, FCallerFile, FCallerLine),
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
		Format:   fmt.Sprintf("%s | %s{%s} | %s | %s:%s", FTimestamp, FColor, FName, FContent, FCallerFile, FCallerLine),
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
		Format:   fmt.Sprintf("%s | %s{%s} | %s | %s:%s", FTimestamp, FColor, FName, FContent, FCallerFile, FCallerLine),
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
		Format:   fmt.Sprintf("%s | %s{%s} | %s | %s:%s", FTimestamp, FColor, FName, FContent, FCallerFile, FCallerLine),
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
		Format:   fmt.Sprintf("%s | %s{%s} | %s | %s:%s", FTimestamp, FColor, FName, FContent, FCallerFile, FCallerLine),
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
