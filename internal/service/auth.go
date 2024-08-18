package service

import (
	"fmt"
	"github.com/AlnurZhanibek/kazusa-server/internal/entity"
	"github.com/AlnurZhanibek/kazusa-server/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(email string, password string) (string, error) {
	users, err := s.userRepo.Read(
		entity.Pagination{
			Offset: 0,
			Limit:  1,
		},
		repository.UserFilters{
			Email: email,
		},
	)
	if err != nil {
		return "", fmt.Errorf("auth service login error: %v", err)
	}
	if len(users) == 0 {
		return "", fmt.Errorf("auth service login error: user not found")
	}
	user := users[0]

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("auth service login error comparing pass: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Subject:   email,
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("auth service login error generating token: %v", err)
	}

	return tokenString, nil
}

func (s *AuthService) Register(name string, email string, phone string, password string, passwordConfirmation string) (string, error) {
	if password != passwordConfirmation {
		return "", fmt.Errorf("auth service register error: passwords do not match")
	}

	if name == "" || email == "" {
		return "", fmt.Errorf("auth service register error: name or email is empty")
	}

	users, err := s.userRepo.Read(
		entity.Pagination{
			Offset: 0,
			Limit:  1,
		},
		repository.UserFilters{
			Email: email,
		},
	)
	if err != nil {
		return "", fmt.Errorf("auth service register error: %v", err)
	}
	if len(users) != 0 {
		return "", fmt.Errorf("auth service register error: user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("auth service register password hash error: %v", err)
	}

	newUser := repository.User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: string(hashedPassword),
		Role:     "user",
	}

	_, err = s.userRepo.Create(newUser)
	if err != nil {
		return "", fmt.Errorf("auth service register error: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Subject:   email,
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("auth service register error generating token: %v", err)
	}

	return tokenString, nil
}
