package init_json

import (
	"encoding/json"
	"fmt"
	json2 "github.com/NoOl01/golog/internal/json"
	"os"
)

func WriteJson(file *os.File) error {
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(file)

	config := json2.Json{
		LogLevels: []json2.LogLevel{
			{
				Name:    "INFO",
				Debug:   false,
				Format:  "$time_stamp | $color[$name] | $content | $caller",
				Color:   "#FFFFFF",
				Outputs: []string{"console", "file"},
				Rotation: json2.Rotation{
					AutoDelete:     true,
					AutoDeleteDays: 10,
					Compress:       true,
					CompressDays:   3,
				},
			},
			{
				Name:    "DEBUG",
				Debug:   true,
				Format:  "$time_stamp | $color[$name] | $content | $caller",
				Color:   "#00FF00",
				Outputs: []string{"console", "file"},
				Rotation: json2.Rotation{
					AutoDelete:     true,
					AutoDeleteDays: 10,
					Compress:       true,
					CompressDays:   3,
				},
			},
			{
				Name:    "ERROR",
				Debug:   false,
				Format:  "$time_stamp | $color[$name] | $content | $caller",
				Color:   "#FF0000",
				Outputs: []string{"console", "file"},
				Rotation: json2.Rotation{
					AutoDelete:     true,
					AutoDeleteDays: 10,
					Compress:       true,
					CompressDays:   3,
				},
			},
		},
		FileConfig: json2.FileConfig{
			FileExtension: "txt",
			FilePath:      "./example",
			FileMaxSize:   10,
		},
	}

	jsonText, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonText)
	if err != nil {
		return err
	}

	return nil
}
