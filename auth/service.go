package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(userID string) (string, string, error)
	ValidateToken(token string) (*jwt.Token, error)
	RefreshToken(refreshToken string) (string, error)
}

type jwtService struct {
	SecretKey        []byte
	RefreshSecretKey []byte
}

func NewService() *jwtService {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	refreshSecretKey := []byte(os.Getenv("REFRESH_SECRET_KEY"))
	return &jwtService{SecretKey: secretKey, RefreshSecretKey: refreshSecretKey}
}

func (s *jwtService) GenerateToken(userID string) (string, string, error) {
	accessClaims := jwt.MapClaims{}
	accessClaims["user_id"] = userID
	accessClaims["exp"] = time.Now().Add(time.Hour).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessSignedToken, err := accessToken.SignedString(s.SecretKey)
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{}
	refreshClaims["user_id"] = userID
	refreshClaims["exp"] = time.Now().Add(30 * 24 * time.Hour).Unix()

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshSignedToken, err := refreshToken.SignedString(s.RefreshSecretKey)
	if err != nil {
		return "", "", err
	}

	return accessSignedToken, refreshSignedToken, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return s.SecretKey, nil
	})
	if err != nil {
		return parsedToken, err
	}

	return parsedToken, nil
}

func (s *jwtService) RefreshToken(refreshToken string) (string, error) {
	parsedToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return s.RefreshSecretKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		if time.Now().After(expirationTime) {
			return "", errors.New("refresh token has expired")
		}
	} else {
		return "", errors.New("invalid refresh token")
	}

	userID := parsedToken.Claims.(jwt.MapClaims)["user_id"].(string)
	accessToken, _, err := s.GenerateToken(userID)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
