package util

import (
	"crypto/hmac"
	"crypto/sha256"
)

func GenerateHMACHash(payload string, key string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(payload))
	return mac.Sum(nil)
}
