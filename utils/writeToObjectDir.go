package utils

import "os"

func WriteToObjectsDir(path string, content []byte, isDir bool)  {
	objectDirFileName, _ := GetSHA1(content)
	// objectDirContent := "blob\n" + utils.EncodeBase64(content)
	objectDirContent := ""
	if(isDir) {
		objectDirContent += string(content)
	} else {
		objectDirContent += "blob\n" + string(content)
	}
	os.MkdirAll(path + "/objects/" + objectDirFileName[0:2], 0755)
	file, err := os.Create(path + "/objects/" + objectDirFileName[0:2] + "/" + objectDirFileName[2:])
	if err != nil {
	  panic(err)
	}
	defer file.Close()
	file.WriteString(objectDirContent)
}