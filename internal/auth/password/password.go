package password

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	return string(hash), err
}

func ComparePassword(
	hashedPassword string,
	password string,
) error {

	return bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
}
