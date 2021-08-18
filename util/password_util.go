package util

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func PasswordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

func GeneratePassword() string {
	return uuid.NewString() + uuid.NewString() + uuid.NewString() + uuid.NewString()
}
