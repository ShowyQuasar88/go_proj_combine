package utils

import (
	"github.com/google/uuid"
	"strings"
)

// NewUUID 生成一个32位的UUID字符串(不含连字符)
func NewUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// NewUUIDWithHyphen 生成一个36位的UUID字符串(含连字符)
func NewUUIDWithHyphen() string {
	return uuid.New().String()
}
