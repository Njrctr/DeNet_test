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

type User interface {
	GetUserInfo(userId int) (models.User, error)
	GetUsersLeaderboard() ([]models.User, error)
}

type Service struct {
	Autorization
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
		User:         NewUserService(repos.User),
	}
}
