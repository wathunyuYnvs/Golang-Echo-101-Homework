package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodePassword(p string) string {
	hasher := md5.New()
	hasher.Write([]byte(p))
	return hex.EncodeToString(hasher.Sum(nil))
}
func ComparePassword(hash_password string, password string) bool {
	return EncodePassword(password) == hash_password
}
