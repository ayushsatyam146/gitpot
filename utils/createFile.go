package utils

import (
	"os"
	"strings"
)

func CreateFile(filePath string, content []byte) string {
	_, err := os.Stat(filePath)
	if err == nil {
		return "File already exists"
	}
	dirPath := strings.Join(strings.Split(filePath, "/")[:len(strings.Split(filePath, "/"))-1], "/")
	CreateDir(dirPath)
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(content)
	return "File created successfully"
}

func CreateDir(dirPath string) string {
	_, err := os.Stat(dirPath)
	if err == nil {
		return "Directory already exists"
	}
	read_err := os.MkdirAll(dirPath, 0755)
	if read_err != nil {
		panic(read_err)
	}
	return "Directory created successfully"
}