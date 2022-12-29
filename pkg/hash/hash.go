package hash

import (
	// "crypto/sha256"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

const salt = "qscftyjmkoplmhtf567tfcxsq"

//With SHA256
// func GeneratePasswordHash(password string) string {
// 	h := sha256.Sum256([]byte(password))

// 	return fmt.Sprintf("%x", h)
// }

//-Other variant with scrypt and salt
func GeneratePasswordHash(password string) string {
	hash, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)

	return fmt.Sprintf("%x", hash)
}
