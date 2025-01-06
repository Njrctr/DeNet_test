package service

import (
	"github.com/Njrctr/DeNet_test/internal/models"
	"github.com/Njrctr/DeNet_test/internal/repository"
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

type Task interface {
	CreateTask(task models.TaskCreate) (int, error)
	CompleteTask(userId, taskId int) error
	GetAllTasks() ([]models.Task, error)
}

type Service struct {
	Autorization
	User
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
		User:         NewUserService(repos.User),
		Task:         NewTaskService(repos.Task),
	}
}
