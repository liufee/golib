package crypt

import (
	"crypto/md5"
	"hash"
	"encoding/hex"
)

func Md5(str string) string {

	var h hash.Hash = md5.New()
	h.Write([]byte(str))
	var by []byte = h.Sum(nil)
	return hex.EncodeToString(by)
}
