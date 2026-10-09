package otp

import (
	"crypto/rand"
	"math/big"
)

func GenerateOTP() (string, error) {

	const otpLength = 6

	otp := ""

	for i := 0; i < otpLength; i++ {

		number, err := rand.Int(
			rand.Reader,
			big.NewInt(10),
		)

		if err != nil {
			return "", err
		}

		otp += number.String()
	}

	return otp, nil
}
