package generate_code

import (
	"github.com/NoOl01/golog/internal/generate_code/json_decode"
	"github.com/NoOl01/golog/internal/json"
)

func GenerateCode(filePath string) error {
	newJson := json.Json{}
	if err := json_decode.JsonDecode(filePath, &newJson); err != nil {
		return err
	}

	return nil
}
