package service

import (
	"github.com/Njrctr/DeNet_test/models"
	"github.com/Njrctr/DeNet_test/pkg/repository"
)

type Autorization interface {
	CreateUser(user models.SignUpInput) (int, error)
	GenerateJWTToken(user models.SignInInput) (string, error)
	ParseJWTToken(accesToken string) (int, error)
}

type Service struct {
	Autorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
	}
}
