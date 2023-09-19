package security

import "golang.org/x/crypto/bcrypt"

// CheckPasswordHash check password with the hash
func CheckPasswordHash(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
