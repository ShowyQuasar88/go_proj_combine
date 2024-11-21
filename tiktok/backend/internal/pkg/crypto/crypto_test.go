package crypto

import (
	"fmt"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	key, err := GenerateAESKey()
	if err != nil {
		panic(err)
	}
	iv, err := GenerateAESIV()
	if err != nil {
		panic(err)
	}

	fmt.Printf("AES Key: %s\n", key)
	fmt.Printf("AES IV: %s\n", iv)
}
