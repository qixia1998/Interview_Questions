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

// Base64UrlEncoding base64编码
func Base64UrlEncoding(str string) string {
	src := []byte(str)
	ret := base64.URLEncoding.EncodeToString(src)
	return ret
}

// Base64UrlDecoding base64解码
func Base64UrlDecoding(str string) (string, error) {

	retByte, err := base64.URLEncoding.DecodeString(str)
	ret := string(retByte)
	return ret, err
}
