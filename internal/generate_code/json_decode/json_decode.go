package json_decode

import (
	decodeJson "encoding/json"
	"fmt"
	"github.com/NoOl01/golog/internal/golog_errs"
	"github.com/NoOl01/golog/internal/json"
	"os"
)

func JsonDecode(filePath string, json *json.Json) error {
	if err := checkJson(filePath); err != nil {
		return err
	}

	jsonText, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return decodeJson.Unmarshal(jsonText, &json)
}

func checkJson(filePath string) error {
	file, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	if file.IsDir() {
		return fmt.Errorf("%w", golog_errs.FileNotExists)
	}

	return nil
}
