package crypto

import (
	"crypto/rand"
	"encoding/hex"
)

// generateRandomKey 生成指定长度的随机密钥
func generateRandomKey(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// GenerateAESKey 生成AES-256密钥(32字节)
func GenerateAESKey() (string, error) {
	return generateRandomKey(32)
}

// GenerateAESIV 生成AES IV(16字节)
func GenerateAESIV() (string, error) {
	return generateRandomKey(16)
}
