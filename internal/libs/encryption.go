package libs

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(value string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CheckHash(hashed string, value string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(value))
	return err == nil
}
