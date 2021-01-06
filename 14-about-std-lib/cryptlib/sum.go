package cryptlib

import (
	"crypto/md5"
	"encoding/hex"
)

// GetStrMD5 计算编码
func GetStrMD5(str string) string {
	r := md5.Sum([]byte(str))
	return hex.EncodeToString(r[:])
}
