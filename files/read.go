package files

import (
	"fmt"
	"os"
	"path/filepath"
)

// ReadFileToString reads the contents of the file at the specified path and returns it as a string
func Read(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file not found: %s", filePath)
		}
		return "", err
	}

	return string(content), nil
}

func ListFilesRecursive(path string) ([]string, error) {
	var files []string

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the path is a regular file
		if !info.IsDir() {
			files = append(files, filePath)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
