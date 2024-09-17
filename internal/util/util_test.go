package util

import "testing"

func TestGenerateHMACHash(t *testing.T) {
	secret := "nihal"
	payload := "meowmeow"
	hash := GenerateHmacSHA256Hash(payload, secret)
	t.Log(hash)
}
