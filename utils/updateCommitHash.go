package utils

import "os"

func ReadFile(filePath string) []byte {
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

func UpdateCommitHashInBranch(commitHash string) {
	headContent := ReadFile("test/.gitpot/HEAD")
	fileName := "test/.gitpot/" + string(headContent)
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(commitHash)
}