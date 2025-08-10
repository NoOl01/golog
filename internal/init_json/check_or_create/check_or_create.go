package check_or_create

import (
	"errors"
	"fmt"
	"github.com/NoOl01/golog/internal/golog_errs"
	"os"
)

func CheckOrCreateFolder(folderPath string) error {
	if folderPath == "" {
		folderPath = "./"
	}

	_, err := os.Stat(folderPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(folderPath, os.ModePerm)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	return nil
}

func CheckOrCreateFile(filePath string) (*os.File, error) {
	if filePath == "" {
		return nil, fmt.Errorf("%w", golog_errs.FileNotExists)
	}

	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			file, err := os.Create(filePath)
			if err != nil {
				return nil, err
			}
			return file, nil
		}
		return nil, err
	}

	return file, nil
}
