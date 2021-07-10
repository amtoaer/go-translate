package utils

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

func CheckEmpty(fileToken, defaultToken string) string {
	if fileToken == "" {
		return defaultToken
	}
	return fileToken
}

func GenerateUUID() string {
	return uuid.NewV4().String()
}

func GenerateYoudaoSign(id, content, salt, time, secret string) string {
	var (
		tmp   []rune = []rune(content)
		input string
	)
	if len(tmp) > 20 {
		input = string(tmp[:10]) + strconv.Itoa(len(tmp)) + string(tmp[len(tmp)-10:])
	} else {
		input = content
	}
	signStr := id + input + salt + time + secret
	return fmt.Sprintf("%x", sha256.Sum256([]byte(signStr)))
}
