package model

import (
	"time"

	"github.com/Yash-programs16/devsync/internal/models"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID uuid.UUID `json:"id"`

	Username string `json:"username"`

	Email string `json:"email"`

	DisplayName string `json:"display_name"`

	AvatarURL *string `json:"avatar_url"`

	Status string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
}

func ToUserResponse(
	user *models.User,
) UserResponse {

	return UserResponse{

		ID: user.ID,

		Username: user.Username,

		Email: user.Email,

		DisplayName: user.DisplayName,

		AvatarURL: user.AvatarURL,

		Status: string(user.Status),

		CreatedAt: user.CreatedAt,
	}
}
