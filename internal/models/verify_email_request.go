package models

type VerifyEmailRequest struct {
	Email string `json:"email"`

	OTP string `json:"otp"`
}
