package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/logs"
	"math/rand"
)

const (
	encryptSecretKey = "1234567891234567"
	paddingSize      = 16

	saltSize = 6
	saltPool = "0123456789abcdefghijklmnopqrstuvwxyz"
)

func GetHashSaltyPassword(password string) (string, string, error) {
	salt := getSalt()
	saltyPassword := GetSaltyPassword(password, salt)
	hashSaltyPassword, err := GetHashString(saltyPassword)
	if err != nil {
		logs.Error(err)
		return "", "", err
	}
	return hashSaltyPassword, salt, nil
}

func GetSaltyPassword(password, salt string) string {
	return password + salt
}

func getSalt() string {
	var str string
	for i := 0; i < saltSize; i++ {
		str += string(saltPool[rand.Intn(len(saltPool))])
	}
	return str
}

func GetHashString(originString string) (string, error) {
	shaer := sha1.New()
	if _, err := shaer.Write([]byte(originString)); err != nil {
		logs.Error(err)
		return "", err
	}
	return fmt.Sprintf("%x", shaer.Sum(nil)), nil
}

func Encrypt(originString string) (string, error) {
	paddingBytes := getPaddingBytes([]byte(originString))
	resultBytes := make([]byte, len(paddingBytes))
	block, err := aes.NewCipher([]byte(encryptSecretKey))
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCEncrypter(block, []byte(encryptSecretKey))
	blockMode.CryptBlocks(resultBytes, paddingBytes)
	return base64.StdEncoding.EncodeToString(resultBytes), nil // 对加密后的结果进行Base64编码
}

func Decrypt(encryptedString string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}
	paddingBytes := make([]byte, len(decodedBytes))
	block, err := aes.NewCipher([]byte(encryptSecretKey))
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(encryptSecretKey))
	blockMode.CryptBlocks(paddingBytes, decodedBytes)
	return string(reducePaddingBytes(paddingBytes)), nil
}

func getPaddingBytes(originBytes []byte) []byte {
	lengthenLength := paddingSize - len(originBytes)%paddingSize
	additionBytes := bytes.Repeat([]byte{byte(lengthenLength)}, lengthenLength)
	return append(originBytes, additionBytes...)
}

func reducePaddingBytes(paddingBytes []byte) []byte {
	if len(paddingBytes) == 0 {
		return paddingBytes
	}
	additionBytesLength := int(paddingBytes[len(paddingBytes)-1])
	return paddingBytes[:len(paddingBytes)-additionBytesLength]
}
