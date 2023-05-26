package encry

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type AesEncrypt struct {
	key   []byte
	iv    []byte
	block cipher.Block
}

func NewAesEncryptInstance(key string, secret string) (*AesEncrypt, error) {
	b, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	a := &AesEncrypt{
		key:   []byte(key),
		iv:    []byte(secret),
		block: b,
	}
	return a, nil
}

// AesBase64Encrypt aes Base64 加密
func (a *AesEncrypt) AesBase64Encrypt(in string) (string, error) {
	origData := []byte(in)
	origData = PKCS5Adding(origData, a.block.BlockSize())
	crypted := make([]byte, len(origData))
	bm := cipher.NewCBCEncrypter(a.block, a.iv)
	bm.CryptBlocks(crypted, origData)
	var b = base64.StdEncoding.EncodeToString(crypted)
	return b, nil
}

// AesBase64Decrypt aes Base64 解密
func (a *AesEncrypt) AesBase64Decrypt(b string) (string, error) {
	crypted, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return "", err
	}
	origData := make([]byte, len(crypted))
	bm := cipher.NewCBCDecrypter(a.block, a.iv)
	bm.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	var out = string(origData)
	return out, nil
}

func PKCS5Adding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
