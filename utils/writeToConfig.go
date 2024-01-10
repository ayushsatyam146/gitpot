package utils

import (
	"fmt"
	"os"
	"strings"
)

func WriteToConfig(key string, value string) {
	fmt.Println("Write to config")
	configFile := "test/.gitpot/config"
	content, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	key = strings.Split(key,".")[1]
	contentString := string(content)
	if strings.Contains(contentString, "[user]") {
		// if key exists, replace value
		if strings.Contains(contentString, key) {
			previousValue := strings.Split(strings.Split(contentString, key + " = ")[1], "\n")[0]
			contentString = strings.ReplaceAll(contentString, key + " = " + previousValue, key + " = " + value)
			os.WriteFile(configFile, []byte(contentString), 0644)
			return
		}
		previousContentString := strings.Split(contentString, "[user]")[0]
		contentString = "[user]\t" + strings.Split(contentString, "[user]")[1]
		contentString+= "\t" + key + " = " + value + "\n"
		os.WriteFile(configFile, []byte(string(previousContentString) + contentString), 0644)
	} else {
		contentString = contentString + "\n" + "[user]\n\t" + key + " = " + value + "\n" 
		os.WriteFile(configFile, []byte(contentString), 0644)
	}
}

func ReadFromConfig(key string) string {
	configFile := "test/.gitpot/config"
	content, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	key = strings.Split(key,".")[1]
	contentString := string(content)
	if strings.Contains(contentString, "[user]") {
		if strings.Contains(contentString, key) {
			value := strings.Split(strings.Split(contentString, key + " = ")[1], "\n")[0]
			// fmt.Println(key, " = " + value)
			return value
		}
	}
	return ""
}