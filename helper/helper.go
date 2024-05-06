package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

// sha256加密
func Sha256cry(password string) string {
	h := sha256.Sum256([]byte(password))
	// 返回16进制字符串
	return hex.EncodeToString(h[:])
}
