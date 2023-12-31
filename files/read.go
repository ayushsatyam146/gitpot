package files

import (
	"os"
	"path/filepath"
)

// ReadFileToString reads the contents of the file at the specified path and returns it as a string
func ReadFile(filePath string) ([]byte) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	// read the file
	fileSize := stat.Size()
	bytes := make([]byte, fileSize)
	_, err = file.Read(bytes)
	if err != nil {
		panic(err)
	}

	return bytes
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
