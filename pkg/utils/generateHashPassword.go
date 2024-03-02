package utils

import "golang.org/x/crypto/bcrypt"

// This takes the user plain password and generate a hash out of user
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
