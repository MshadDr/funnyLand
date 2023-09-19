package security

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword make a hash from a password
func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}

func EncryptWhatsappToken(userId int, phone string) (string, error) {
	tokenString := fmt.Sprintf("%d@%s", userId, phone)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(tokenString), 14)
	return string(hashedPassword), err
}
