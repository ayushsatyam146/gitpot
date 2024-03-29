package utils

import (
	"os"
)

func WriteToObjectsDir(path string, content []byte, notFile bool) string {
	objectDirFileName, _ := GetSHA1(content)

	objectDirContent := ""
	if notFile {
		objectDirContent += string(content)
	} else {
		objectDirContent += "blob\n" + string(content)
	}
	// objectDirContent = EncodeBase64([]byte(objectDirContent))
	os.MkdirAll(path+"/objects/"+objectDirFileName[0:2], 0755)
	file, err := os.Create(path + "/objects/" + objectDirFileName[0:2] + "/" + objectDirFileName[2:])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(objectDirContent)

	return objectDirFileName
}
