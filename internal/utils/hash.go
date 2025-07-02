package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
)

func GenerateHash(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func CompareHash(hash1, hash2 string) bool {
	return subtle.ConstantTimeCompare([]byte(hash1), []byte(hash2)) == 1
}
