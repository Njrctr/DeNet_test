package service

import (
	"github.com/Njrctr/DeNet_test/internal/models"
	"github.com/Njrctr/DeNet_test/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) GetUserInfo(userId int) (models.User, error) {
	return u.repo.GetUserInfo(userId)
}

func (u *UserService) GetUsersLeaderboard() ([]models.User, error) {
	return u.repo.GetUsersLeaderboard()
}
