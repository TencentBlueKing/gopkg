package stringx

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hash calculate the md5 of string
func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
