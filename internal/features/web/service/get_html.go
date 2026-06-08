package web_service

import (
	"fmt"
	"os"
	"path"
)

func (w *WebService) GetFile() ([]byte, error) {
	filePath := path.Join(os.Getenv("PROJECT_DIR"), "/public/index.html")

	html, err := w.webRepository.GetFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("get file from repository: %w", err)
	}

	return html, nil

}
