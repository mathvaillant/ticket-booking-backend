package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mathvaillant/ticket-booking-project-v0/models"
	"github.com/mathvaillant/ticket-booking-project-v0/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repository models.AuthRepository
}

func (s *AuthService) Login(ctx context.Context, loginData *models.AuthCredentials) (string, *models.User, error) {
	user, err := s.repository.GetUser(ctx, "email = ?", loginData.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, fmt.Errorf("invalid credentials")
		}
		return "", nil, err
	}

	if !models.MatchesHash(loginData.Password, user.Password) {
		return "", nil, fmt.Errorf("invalid credentials")
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 168).Unix(),
	}

	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, os.Getenv("JWT_SECRET"))

	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *AuthService) Register(ctx context.Context, registerData *models.AuthCredentials) (string, *models.User, error) {
	if !models.IsValidEmail(registerData.Email) {
		return "", nil, fmt.Errorf("please, provide a valid email to register")
	}

	if _, err := s.repository.GetUser(ctx, "email = ?", registerData.Email); !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", nil, fmt.Errorf("the user email is already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil, err
	}

	registerData.Password = string(hashedPassword)

	user, err := s.repository.RegisterUser(ctx, registerData)
	if err != nil {
		return "", nil, err
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 168).Unix(),
	}

	// Generate the JWT
	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, os.Getenv("JWT_SECRET"))
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func NewAuthService(repository models.AuthRepository) models.AuthService {
	return &AuthService{
		repository: repository,
	}
}
