package postgres

import (
	"fmt"

	"github.com/Njrctr/DeNet_test/internal/models"
	"github.com/Njrctr/DeNet_test/pkg/utils"
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
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, refer_code) VALUES ($1, $2, $3) RETURNING id", usersTable)

	referCode := utils.RandStringBytes()
	result := a.db.QueryRow(query, user.Username, user.Password, referCode)
	if err := result.Scan(&userId); err != nil {
		value, ok := customErrors[err.Error()]
		if ok {
			return 0, fmt.Errorf("%s", value)
		}

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
