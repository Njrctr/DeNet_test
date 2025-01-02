package postgres

import (
	"fmt"

	"github.com/Njrctr/DeNet_test/models"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user models.SignUpInput) (int, error) {
	var userId int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) VALUES ($1, $2) RETURNING id", usersTable)

	result := a.db.QueryRow(query, user.Username, user.Password)
	if err := result.Scan(&userId); err != nil {
		return 0, err
	}

	return userId, nil
}

func (a *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := a.db.Get(&user, query, username, password)

	return user, err
}
