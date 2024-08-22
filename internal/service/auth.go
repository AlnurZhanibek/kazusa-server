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

type Claims struct {
	jwt.StandardClaims
	Name string      `json:"name"`
	Role entity.Role `json:"role"`
}

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
		entity.UserFilters{
			Email: &email,
		},
	)
	if err != nil {
		return "", fmt.Errorf("auth service login error: %v", err)
	}
	if len(users) == 0 {
		return "", fmt.Errorf("auth service login error: user not found")
	}
	user := users[0]

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("auth service login error comparing pass: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Subject:   email,
		},
		Name: user.Name,
		Role: user.Role,
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
		entity.UserFilters{
			Email: &email,
		},
	)
	if err != nil {
		return "", fmt.Errorf("auth service register error: %v", err)
	}
	if len(users) != 0 {
		return "", fmt.Errorf("auth service register error: user already exists")
	}

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("auth service register password hash error: %v", err)
	}

	hashedPassword := string(hashedPasswordBytes)

	newUser := entity.NewUser{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: hashedPassword,
	}

	_, err = s.userRepo.Create(newUser)
	if err != nil {
		return "", fmt.Errorf("auth service register error: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Subject:   email,
		},
		Name: name,
		Role: "user",
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("auth service register error generating token: %v", err)
	}

	return tokenString, nil
}
