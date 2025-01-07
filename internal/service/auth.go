package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Njrctr/DeNet_test/internal/models"
	"github.com/Njrctr/DeNet_test/internal/repository"
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
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
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
	token, err := jwt.ParseWithClaims(accesToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
