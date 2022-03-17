package rc4

import (
	"testing"
)

const (
	key = "asdfasdf23a4"
)

func TestEncrypt(t *testing.T) {
	t.Log(New(key).Encrypt("123456"))
}

func TestDecrypt(t *testing.T) {
	t.Log(New(key).Decrypt("0bd9ca20c591"))
}
