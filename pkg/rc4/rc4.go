package rc4

import (
	cryptoRc4 "crypto/rc4"
	"encoding/hex"
	"fmt"
)

var _ Rc4 = (*rc4)(nil)

type Rc4 interface {
	i()
	// Encrypt 加密
	Encrypt(encryptStr string) (string, error)

	// Decrypt 解密
	Decrypt(decryptStr string) (string, error)
}

type rc4 struct {
	key string
}

func New(key string) Rc4 {
	return &rc4{
		key: key,
	}
}

func (r *rc4) i() {}

// Encrypt implements Rc4
func (r *rc4) Encrypt(encryptStr string) (string, error) {
	dest1 := make([]byte, len(encryptStr))
	cipher1, err := cryptoRc4.NewCipher([]byte(r.key))
	if err != nil {
		return "", err
	}

	cipher1.XORKeyStream(dest1, []byte(encryptStr))
	fmt.Println(string([]byte(r.key)))
	return hex.EncodeToString(dest1), nil
}

// Decrypt implements Rc4
func (r *rc4) Decrypt(decryptStr string) (string, error) {
	hexDecryptStr := make([]byte, len(decryptStr))
	dest2, _ := hex.DecodeString(decryptStr)
	cipher2, err := cryptoRc4.NewCipher([]byte(r.key))
	if err != nil {
		return "", err
	}

	cipher2.XORKeyStream(hexDecryptStr, dest2)
	return string(hexDecryptStr), nil
}
