package hash

import (
	// "crypto/sha256"
	"fmt"

	"golang.org/x/crypto/scrypt"
)

//With SHA256
// func GeneratePasswordHash(password, salt string) string {
// 	h := sha256.Sum256([]byte(password))

// 	return fmt.Sprintf("%x", h)
// }

//-Other variant with scrypt and salt
func GeneratePasswordHash(password, salt string) string {
	// salt := os.Getenv("SALT") other variant getting salt from .env
	hash, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)

	return fmt.Sprintf("%x", hash)
}
