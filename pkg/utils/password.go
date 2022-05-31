package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(hash, input string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(input)) == nil
}
