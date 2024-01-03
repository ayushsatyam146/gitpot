package files

import "os"

func Write(filePath string, content []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(content)
	return nil
}
