package user

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	FindById(id string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) Register(input RegisterUserInput) (User, error) {
	user := User{
		ID:        uuid.New(),
		Username:  input.Username,
		Email:     input.Email,
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	saltBytes := make([]byte, 16)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return user, err
	}
	salt := hex.EncodeToString(saltBytes)

	passwordHash, err := bcrypt.GenerateFromPassword(append([]byte(input.Password), saltBytes...), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Salt = salt

	newUser, err := s.repository.Create(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == uuid.Nil {
		return user, errors.New("email or password is incorrect")
	}

	saltBytes, err := hex.DecodeString(user.Salt)
	if err != nil {
		return User{}, err
	}

	combinedPassword := append([]byte(password), saltBytes...)
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), combinedPassword)
	if err != nil {
		return User{}, errors.New("email or password is incorrect")
	}

	return user, nil
}

func (s *service) FindById(id string) (User, error) {
	user, err := s.repository.FindByID(id)

	if err != nil {
		return user, err
	}

	return user, nil
}
