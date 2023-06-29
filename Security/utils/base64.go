package utils

import "encoding/base64"

// Base64Encoding base64编码
func Base64Encoding(str string) string {
	src := []byte(str)
	ret := base64.StdEncoding.EncodeToString(src)
	return ret
}

// Base64Decoding base64解码
func Base64Decoding(str string) (string, error) {

	retByte, err := base64.StdEncoding.DecodeString(str)
	ret := string(retByte)
	return ret, err
}
