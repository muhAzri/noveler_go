package user

import (
	"github.com/google/uuid"
)

type UserFormatter struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

func FormatUser(user User, accessToken string, refreshToken string) UserFormatter {
	formatter := UserFormatter{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return formatter
}
