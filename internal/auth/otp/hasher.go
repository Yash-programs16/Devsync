package otp

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashOTP(code string) string {
	hash := sha256.Sum256([]byte(code))

	return hex.EncodeToString(hash[:])
}

func CompareOTP(code string, hash string) bool {
	return HashOTP(code) == hash
}
