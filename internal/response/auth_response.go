package response

import (
	"github.com/salawatbro/chat-app/internal/models"
	"github.com/salawatbro/chat-app/pkg/constants"
	"time"
)

type RegisterResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

func NewRegisterResponse(user *models.User) *RegisterResponse {
	return &RegisterResponse{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}
}

func NewLoginResponse(token string, expiresAt time.Time) *LoginResponse {
	return &LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt.Format(constants.TimestampFormat),
	}
}
