package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func Bcrypt(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), 14)
	return string(bytes), err
}
