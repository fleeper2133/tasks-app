package service

import (
	"github.com/fleeper2133/tasks-app/internal/domain"
	"github.com/fleeper2133/tasks-app/internal/pkg"
	"github.com/fleeper2133/tasks-app/internal/repository"
)

type Authorization interface {
	CreateUser(user domain.SignUp) (int, error)
	GenerateTokens(input domain.SignIn) (pkg.TokenJWT, error)
	ParseToken(token string) (string, error)
	RefreshToken(refreshToken string) (pkg.TokenJWT, error)
}

type Tasks interface {
	GetAll(userId int) ([]domain.Task, error)
	Create(task domain.TaskInput, userId int) (int, error)
	GetById(taskId int, userId int) (domain.Task, error)
	Delete(taskId int, userId int) error
	Update(taskId int, input domain.TaskUpdate, userId int) error
}

type Service struct {
	Authorization
	Tasks
}

func NewService(repo *repository.Repository, jwtManager *pkg.TokenJWTManager) *Service {
	return &Service{
		Authorization: NewAuthorizationService(repo.Authorization, *jwtManager),
		Tasks:         NewTasksService(repo.Tasks),
	}
}
