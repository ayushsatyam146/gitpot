package index

import (
	"os"
	"strings"
)

func WriteContentToIndex(path string, content []byte) {
	path = path + "/index"
	stringContent := string(content)
	element := strings.Split(stringContent, "\n")[1:]
	stringContent = strings.Join(element, "\n")
	bytes := []byte(stringContent)

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write(bytes)
}
