// ------------------------------
// 7. utils/hash.go
// ------------------------------

package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pw string) (string, error) {
	return bcrypt.GenerateFromPassword([]byte(pw), 14)
}

func CheckPassword(hashed string, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}