package utils

import (
	"bytes"
	"fmt"
	"crypto/md5"
	"log"
	"crypto/aes"
	"encoding/hex"
	"encoding/base64"
)

const (
	AESKEY = "6368616e676520746869732070617373776f726420746f206120736563726574"
)

// 计算 md5 摘要
func MakeMD5(s string) string {
	var buffer *bytes.Buffer = bytes.NewBufferString("")

	if _, error := fmt.Fprintf(buffer, "%x", md5.Sum([]byte(s))); error != nil {
		log.Println(error)
	}

	return buffer.String()
}

// AES ECB 加密
func AESECBEncode(keyString string, plainText string) string {

	key, _ := hex.DecodeString(keyString)

	block, error := aes.NewCipher(key)

	if error != nil {
		log.Fatal(error)
	}

	source := PKCS5Padding([]byte(plainText), block.BlockSize())

	dist := make([]byte, len(source))

	for i := 0; i < len(source); i += block.BlockSize() {
		block.Encrypt(dist[i: i + block.BlockSize()], source[i: i + block.BlockSize()])
	}

	return base64.StdEncoding.EncodeToString([]byte(dist))
}

// AES ECB 解密
func AESECBDecode(keyString string, secretText string) string {

	key, _ := hex.DecodeString(keyString)

	block, error := aes.NewCipher(key)

	if error != nil {
		log.Fatal(error)
	}

	source, error := base64.StdEncoding.DecodeString(secretText)

	if error != nil {
		log.Fatal(error)
	}

	dist := make([]byte, len(source))

	for i := 0; i < len(source); i += block.BlockSize() {
		block.Decrypt(dist[i: i + block.BlockSize()], source[i: i + block.BlockSize()])
	}

	return string(PKCS5UnPadding(dist))
}

func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText) % blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipherText, padText...)
}

func PKCS5UnPadding(originData []byte) []byte {
	length := len(originData)
	unpadding := int(originData[length - 1])

	return originData[:(length - unpadding)]
}

