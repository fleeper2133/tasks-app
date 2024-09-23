package repository

import (
	"fmt"

	"github.com/fleeper2133/tasks-app/internal/domain"
	"github.com/jmoiron/sqlx"
)

type AuthorizationPostgres struct {
	db *sqlx.DB
}

func NewAuthorizationPostgres(db *sqlx.DB) *AuthorizationPostgres {
	return &AuthorizationPostgres{db: db}
}

func (r *AuthorizationPostgres) CreateUser(user domain.SignUp) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, password_hash) VALUES ($1, $2) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthorizationPostgres) GetUser(input domain.SignIn) (domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 and password_hash=$2", userTable)
	if err := r.db.Get(&user, query, input.Email, input.Password); err != nil {
		return domain.User{}, err
	}
	return user, nil

}
