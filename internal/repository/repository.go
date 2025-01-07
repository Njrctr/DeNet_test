package repository

import (
	"github.com/Njrctr/DeNet_test/internal/models"
	"github.com/Njrctr/DeNet_test/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user models.SignUpInput) (int, error)
	GetUser(username, password string) (models.User, error)
}

type User interface {
	GetUserInfo(userId int) (models.User, error)
	GetUsersLeaderboard() ([]models.User, error)
}

type Task interface {
	CreateTask(models.TaskCreate) (int, error)
	CompleteTask(userId, taskId int) error
	GetAllTasks() ([]models.Task, error)
	ReferrerCode(userId int, refCode string) error
}

type Repository struct {
	Autorization
	User
	Task
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: postgres.NewAuthPostgres(db),
		User:         postgres.NewUserPostgres(db),
		Task:         postgres.NewTaskPostgres(db),
	}
}
