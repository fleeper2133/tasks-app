package service

import (
	"github.com/fleeper2133/tasks-app/internal/domain"
	"github.com/fleeper2133/tasks-app/internal/repository"
)

type TasksService struct {
	repo repository.Tasks
}

func NewTasksService(repo repository.Tasks) *TasksService {
	return &TasksService{repo: repo}
}

func (s *TasksService) GetAll(userId int) ([]domain.Task, error) {
	return s.repo.GetAll(userId)
}

func (s *TasksService) Create(task domain.TaskInput, userId int) (int, error) {
	return s.repo.Create(task, userId)
}

func (s *TasksService) GetById(taskId int, userId int) (domain.Task, error) {
	return s.repo.GetById(taskId, userId)
}

func (s *TasksService) Delete(taskId int, userId int) error {
	return s.repo.Delete(taskId, userId)
}

func (s *TasksService) Update(taskId int, input domain.TaskUpdate, userId int) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(taskId, input, userId)
}
