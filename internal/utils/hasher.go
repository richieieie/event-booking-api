package utils

import "golang.org/x/crypto/bcrypt"

const hashingCost = bcrypt.DefaultCost

func HashPassword(password string) (string, error) {
	data, err := bcrypt.GenerateFromPassword([]byte(password), hashingCost)

	return string(data), err
}

func CheckHashedPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
