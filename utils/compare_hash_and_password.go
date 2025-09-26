package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func CompareHashAndPassword(hashedPassword, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
