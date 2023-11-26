package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// lower case
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// upper case
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// encrypt password
func MakePassword(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

// decrypt password
func ValidatePassword(plainpwd, salt, password string) bool {
	return MakePassword(plainpwd, salt) == password
}
