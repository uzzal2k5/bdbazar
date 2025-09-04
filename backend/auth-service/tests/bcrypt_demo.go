package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	storedHash := "$2a$10$SISVYl.X/icxymCrCesFDu6b3QJ/qlz22IczAgSJNbbvLtTrt0c4S"
	password := "strongpassword123"

	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		fmt.Println("❌ Password mismatch:", err)
	} else {
		fmt.Println("✅ Password matched")
	}
}
