package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"testing"
)

// TestGenerateSecretKey 生成32位的密钥
func TestGenerateSecretKey(t *testing.T) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	fmt.Println(base64.URLEncoding.EncodeToString(bytes), nil)
}
