package init_json

import (
	"github.com/NoOl01/golog/internal/init_json/check_or_create"
	path "path/filepath"
)

func InitJsonConfig(filePath string) error {
	dirPath := path.Dir(filePath)

	err := check_or_create.CheckOrCreateFolder(dirPath)
	if err != nil {
		return err
	}

	file, err := check_or_create.CheckOrCreateFile(filePath)
	if err != nil {
		return err
	}

	err = WriteJson(file)
	if err != nil {
		return err
	}

	return nil
}
