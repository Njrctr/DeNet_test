package repository

import (
	"github.com/Njrctr/DeNet_test/models"
	"github.com/Njrctr/DeNet_test/pkg/repository/postgres"
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

type Repository struct {
	Autorization
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: postgres.NewAuthPostgres(db),
		User:         postgres.NewUserPostgres(db),
	}
}
