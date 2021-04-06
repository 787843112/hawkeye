package util

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5V ...
func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
