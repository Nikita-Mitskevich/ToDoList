package web_repository

import (
	"fmt"
	"os"
	core_errors "restapi/internal/core/errors"
)

func (w *WebRepository) GetFile(FilePath string) ([]byte, error) {
	html, err := os.ReadFile(FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file: %s: %w", FilePath, core_errors.ErrNotFound)
		}
		return nil, fmt.Errorf("get file: %s: %w", FilePath, err)
	}

	return html, nil
}
