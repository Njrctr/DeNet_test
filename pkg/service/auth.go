package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Njrctr/DeNet_test/models"
	"github.com/Njrctr/DeNet_test/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "jh23hriuhq9ah982333zvr234fwe"
	signingKey = "fsdglk;asd3j41lqwwkq23xvbxcvbxvcxvca"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.SignUpInput) (int, error) {
	return 0, nil // todo CREATE USER
}

func (s *AuthService) GenerateJWTToken(userInput models.SignInInput) (string, error) {
	user, err := s.repo.GetUser(userInput.Username, generatePasswordHash(userInput.Password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseJWTToken(accesToken string) (int, error) {
	return 0, nil // todo PARSE JWT
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}