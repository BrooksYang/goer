package aes

import (
	"crypto/aes"
	"crypto/cipher"
	b64 "encoding/base64"
)

type AesClient struct {
	Key string
	Iv  string
}

func (ae *AesClient) CFBEncrypt(text string) string {
	iv := []byte(ae.Iv)

	iva := iv[:]
	block, err := aes.NewCipher([]byte(ae.Key))
	if err != nil {
		panic(err)
	}

	byteText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iva)
	encrypted := make([]byte, len(byteText))
	cfb.XORKeyStream(encrypted, byteText)

	return b64.StdEncoding.EncodeToString(encrypted)
}

func (ae *AesClient) CFBDecrypt(encrypted string) string {
	ct, _ := b64.StdEncoding.DecodeString(encrypted)

	iv := []byte(ae.Iv)

	iva := iv[:]

	block, err := aes.NewCipher([]byte(ae.Key))
	if err != nil {
		panic(err)
	}

	cfb := cipher.NewCFBDecrypter(block, iva)
	dst := make([]byte, len(ct))

	cfb.XORKeyStream(dst, ct)

	return string(dst)
}
