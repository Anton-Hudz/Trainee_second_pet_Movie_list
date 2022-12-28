package hash

import (
	// "crypto/sha256"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

const salt = "qscftyjmkoplmhtf567tfcxsq"

func GeneratePasswordHash(password string) (string, error) {

	s := []byte(salt)
	hash, err := scrypt.Key([]byte(password), s, 5, 19, 25, 16)
	if err != nil {
		return "", fmt.Errorf("error while hashing password: %w", err)
	}

	return fmt.Sprintf("%x", hash), nil
}

// func GeneratePasswordHash(password string) string {
// 	h := sha256.Sum256([]byte(someString))

// 	return fmt.Sprintf("%x", h)
// }
