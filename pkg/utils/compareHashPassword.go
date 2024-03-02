package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// This function takes the user password and check against hash stored in the database
func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
