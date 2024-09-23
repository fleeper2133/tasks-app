package repository

import (
	"github.com/fleeper2133/tasks-app/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user domain.SignUp) (int, error)
	GetUser(input domain.SignIn) (domain.User, error)
}

type Tasks interface {
	GetAll(userId int) ([]domain.Task, error)
	Create(task domain.TaskInput, userId int) (int, error)
	GetById(taskId int, userId int) (domain.Task, error)
	Delete(taskId int, userId int) error
	Update(taskId int, input domain.TaskUpdate, userId int) error
}

type Repository struct {
	Authorization
	Tasks
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthorizationPostgres(db),
		Tasks:         NewTasksPostgres(db),
	}
}
