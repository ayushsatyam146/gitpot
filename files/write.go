package files

import "os"

// WriteToFileOverwrite writes content to a file at the specified path, overwriting the file if it exists
func Write(filePath string, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
