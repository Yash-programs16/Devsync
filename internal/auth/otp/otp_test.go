package otp_test

import (
	"testing"

	"github.com/Yash-programs16/devsync/internal/auth/otp"
)

func TestOTP(t *testing.T) {

	code, err := otp.GenerateOTP()

	if err != nil {
		t.Fatal(err)
	}

	hash := otp.HashOTP(code)

	if !otp.CompareOTP(code, hash) {
		t.Fatal("otp verification failed")
	}
}
