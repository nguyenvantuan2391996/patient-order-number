package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func EncodePasswordSHA1(str string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(str))
	return hex.EncodeToString(algorithm.Sum(nil))
}
