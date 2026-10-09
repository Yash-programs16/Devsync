package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Yash-programs16/devsync/internal/auth/otp"
	"github.com/Yash-programs16/devsync/internal/auth/password"
	"github.com/Yash-programs16/devsync/internal/auth/repository"
	"github.com/Yash-programs16/devsync/internal/email"
	"github.com/Yash-programs16/devsync/internal/models"
	"github.com/Yash-programs16/devsync/internal/token"
	"github.com/google/uuid"
)

type AuthService struct {
	userRepo *repository.UserRepository

	refreshTokenRepo *repository.RefreshTokenRepository

	emailVerificationRepo *repository.EmailVerificationRepository

	tokenMaker token.Maker

	emailSender email.Sender
}

func NewAuthService(
	userRepo *repository.UserRepository,

	refreshTokenRepo *repository.RefreshTokenRepository,

	emailVerificationRepo *repository.EmailVerificationRepository,

	tokenMaker token.Maker,

	emailSender email.Sender,
) *AuthService {

	return &AuthService{

		userRepo: userRepo,

		refreshTokenRepo: refreshTokenRepo,

		emailVerificationRepo: emailVerificationRepo,

		tokenMaker: tokenMaker,

		emailSender: emailSender,
	}
}

// Register creates a new user
func (s *AuthService) Register(
	ctx context.Context,
	req models.RegisterRequest,
) (*models.User, error) {

	// 1. Check if email already exists

	existingUser, err := s.userRepo.GetUserByEmail(
		ctx,
		req.Email,
	)

	if err == nil && existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// 2. Check if username already exists

	existingUser, err = s.userRepo.GetUserByUsername(
		ctx,
		req.Username,
	)

	if err == nil && existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// 3. Hash password

	passwordHash, err := password.HashPassword(
		req.Password,
	)

	if err != nil {
		return nil, err
	}

	// 4. Create User model

	user := &models.User{

		ID: uuid.New(),

		Username: req.Username,

		Email: req.Email,

		PasswordHash: passwordHash,

		DisplayName: req.DisplayName,
	}

	// 5. Save user in database

	err = s.userRepo.CreateUser(
		ctx,
		user,
	)

	if err != nil {
		return nil, err
	}

	otpCode, err := otp.GenerateOTP()

	if err != nil {
		return nil, err
	}
	fmt.Println("OTP:", otpCode)

	otpHash := otp.HashOTP(
		otpCode,
	)

	verification := &models.EmailVerification{

		ID: uuid.New(),

		UserID: user.ID,

		Email: user.Email,

		OTPHash: otpHash,

		ExpiresAt: time.Now().Add(
			10 * time.Minute,
		),

		CreatedAt: time.Now(),
	}

	err = s.emailVerificationRepo.Create(
		ctx,
		verification,
	)

	if err != nil {
		return nil, err
	}

	// Send OTP email

	err = s.emailSender.SendEmail(
		user.Email,
		"Verify your DevSync account",
		"Your verification OTP is: "+otpCode,
	)

	if err != nil {
		return nil, err
	}

	// 6. Return created user

	return user, nil
}

func (s *AuthService) Login(
	ctx context.Context,
	req models.LoginRequest,
) (*models.AuthResponse, error) {

	// 1. Find user by email

	user, err := s.userRepo.GetUserByEmail(
		ctx,
		req.Email,
	)

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.EmailVerified {
		return nil, errors.New(
			"please verify your email first",
		)
	}
	// 2. Compare password

	err = password.ComparePassword(
		user.PasswordHash,
		req.Password,
	)

	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// 3. Create access token

	accessToken, err := s.tokenMaker.CreateAccessToken(
		user.ID,
		user.Username,
	)

	if err != nil {
		return nil, err
	}

	// 4. Create refresh token

	refreshToken, err := s.tokenMaker.CreateRefreshToken(
		user.ID,
	)

	if err != nil {
		return nil, err
	}

	// 5. Hash refresh token

	refreshTokenHash := password.HashToken(
		refreshToken,
	)

	// 6. Store refresh token

	refreshTokenModel := &models.RefreshToken{

		ID: uuid.New(),

		UserID: user.ID,

		TokenHash: refreshTokenHash,

		ExpiresAt: time.Now().Add(
			7 * 24 * time.Hour,
		),

		CreatedAt: time.Now(),

		Revoked: false,
	}

	err = s.refreshTokenRepo.CreateRefreshToken(
		ctx,
		refreshTokenModel,
	)

	if err != nil {
		return nil, err
	}

	// 7. Return response

	response := &models.AuthResponse{

		User: models.UserResponse{

			ID: user.ID,

			Username: user.Username,

			Email: user.Email,

			DisplayName: user.DisplayName,

			AvatarURL: user.AvatarURL,

			Status: string(user.Status),

			CreatedAt: user.CreatedAt,
		},

		AccessToken: accessToken,

		RefreshToken: refreshToken,
	}

	return response, nil
}
func (s *AuthService) Refresh(
	ctx context.Context,
	refreshToken string,
) (*models.AuthResponse, error) {

	// 1. Hash incoming refresh token

	tokenHash := password.HashToken(
		refreshToken,
	)

	// 2. Find token in database

	storedToken, err := s.refreshTokenRepo.GetByHash(
		ctx,
		tokenHash,
	)

	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// 3. Check revoked

	if storedToken.Revoked {
		return nil, errors.New("refresh token revoked")
	}

	// 4. Check expiry

	if time.Now().After(
		storedToken.ExpiresAt,
	) {
		return nil, errors.New("refresh token expired")
	}

	// 5. Get user

	user, err := s.userRepo.GetUserByID(
		ctx,
		storedToken.UserID,
	)

	if err != nil {
		return nil, err
	}

	// 6. Create new access token

	accessToken, err := s.tokenMaker.CreateAccessToken(
		user.ID,
		user.Username,
	)

	if err != nil {
		return nil, err
	}

	// 7. Return new access token

	return &models.AuthResponse{

		User: models.UserResponse{

			ID: user.ID,

			Username: user.Username,

			Email: user.Email,

			DisplayName: user.DisplayName,

			AvatarURL: user.AvatarURL,

			Status: string(user.Status),

			CreatedAt: user.CreatedAt,
		},

		AccessToken: accessToken,

		RefreshToken: refreshToken,
	}, nil
}
func (s *AuthService) Logout(
	ctx context.Context,
	refreshToken string,
) error {

	// 1. Hash refresh token

	tokenHash := password.HashToken(
		refreshToken,
	)

	// 2. Revoke token

	err := s.refreshTokenRepo.RevokeToken(
		ctx,
		tokenHash,
	)

	return err
}
func (s *AuthService) VerifyEmail(
	ctx context.Context,
	req models.VerifyEmailRequest,
) error {

	// 1. Find latest OTP record

	verification, err := s.emailVerificationRepo.GetByEmail(
		ctx,
		req.Email,
	)

	if err != nil {
		return errors.New("verification code not found")
	}

	// 2. Check already verified

	if verification.Verified {

		return errors.New("email already verified")

	}

	// 3. Check expiry

	if time.Now().After(
		verification.ExpiresAt,
	) {

		return errors.New("otp expired")

	}

	// 4. Compare OTP

	if !otp.CompareOTP(
		req.OTP,
		verification.OTPHash,
	) {

		return errors.New("invalid otp")

	}

	// 5. Mark verification complete

	err = s.emailVerificationRepo.MarkVerified(
		ctx,
		verification.ID,
	)

	if err != nil {
		return err
	}

	// 6. Update user

	err = s.userRepo.MarkEmailVerified(
		ctx,
		verification.UserID,
	)

	return err
}
