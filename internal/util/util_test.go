package util

import "testing"

func TestGenerateHMACHash(t *testing.T) {
	secret := "nihal"
	payload := "meowmeow"
	hash := GenerateHMACHash(payload, secret)
	t.Log(hash)
}
