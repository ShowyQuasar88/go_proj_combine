package crypto

import (
	"backend/internal/conf"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Crypto struct {
	aesKey     []byte
	aesIV      []byte
	bcryptCost int
}

func NewCrypto(c *conf.Security) (*Crypto, error) {
	// 将十六进制字符串转换为字节数组
	aesKey, err := hex.DecodeString(c.AesKey)
	if err != nil {
		return nil, fmt.Errorf("invalid AES key format: %v", err)
	}

	aesIV, err := hex.DecodeString(c.AesIv)
	if err != nil {
		return nil, fmt.Errorf("invalid AES IV format: %v", err)
	}

	// 验证密钥长度
	if len(aesKey) != 32 { // AES-256需要32字节密钥
		return nil, fmt.Errorf("invalid AES key length: expected 32 bytes, got %d", len(aesKey))
	}

	if len(aesIV) != 16 { // AES需要16字节IV
		return nil, fmt.Errorf("invalid AES IV length: expected 16 bytes, got %d", len(aesIV))
	}

	return &Crypto{
		aesKey:     aesKey,
		aesIV:      aesIV,
		bcryptCost: int(c.BcryptCost),
	}, nil
}

// HashPassword 使用 bcrypt 对密码进行加密
func (c *Crypto) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), c.bcryptCost)
	return string(bytes), err
}

// CheckPasswordHash 验证密码
func (c *Crypto) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Encrypt 使用 AES-CBC 模式加密
func (c *Crypto) Encrypt(text string) (string, error) {
	block, err := aes.NewCipher(c.aesKey)
	if err != nil {
		return "", err
	}

	// 对明文进行PKCS7填充
	plaintext := pkcs7Padding([]byte(text), aes.BlockSize)
	ciphertext := make([]byte, len(plaintext))

	mode := cipher.NewCBCEncrypter(block, c.aesIV)
	mode.CryptBlocks(ciphertext, plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 使用 AES-CBC 模式解密
func (c *Crypto) Decrypt(cryptoText string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(c.aesKey)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("密文太短")
	}

	mode := cipher.NewCBCDecrypter(block, c.aesIV)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// 去除PKCS7填充
	unpaddedPlaintext, err := pkcs7Unpadding(plaintext)
	if err != nil {
		return "", err
	}

	return string(unpaddedPlaintext), nil
}

// pkcs7Padding 实现PKCS7填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// pkcs7Unpadding 实现PKCS7去填充
func pkcs7Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("数据长度为0")
	}

	padding := int(data[length-1])
	if padding > length {
		return nil, errors.New("填充长度错误")
	}

	return data[:length-padding], nil
}
