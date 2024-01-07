package utils

import (
	"os"
)

func CreateFile(filePath string) string {
	_, err := os.Stat(filePath)
	if err == nil {
		return "File already exists"
	}
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	return "File created successfully"
}

func CreateDir(dirPath string) string {
	_, err := os.Stat(dirPath)
	if err == nil {
		return "Directory already exists"
	}
	read_err := os.Mkdir(dirPath, 0755)
	if read_err != nil {
		panic(read_err)
	}
	return "Directory created successfully"
}