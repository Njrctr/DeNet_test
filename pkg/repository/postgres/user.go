package postgres

import (
	"fmt"

	"github.com/Njrctr/DeNet_test/models"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetUserInfo(userId int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id, username, balance, refer_code, refer_from FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, userId)

	return user, err
}

func (r *UserPostgres) GetUsersLeaderboard() ([]models.User, error) {
	var lists []models.User

	query := fmt.Sprintf("SELECT id, username, balance FROM %s ORDER BY balance DESC", usersTable)
	err := r.db.Select(&lists, query)

	return lists, err
}
