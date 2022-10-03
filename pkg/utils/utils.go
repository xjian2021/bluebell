package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(data string) string {
	hash := md5.New()
	hash.Write(nil)
	return hex.EncodeToString(hash.Sum([]byte(data)))
}
