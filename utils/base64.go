package utils

import (
	"encoding/base64"
)

func EncodeBase64(data []byte) (string) {
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded
}

func DecodeBase64(encoded string) ([]byte) {
	decoded,_ := base64.StdEncoding.DecodeString(encoded)
	return decoded
}