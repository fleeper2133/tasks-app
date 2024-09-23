package repository

import (
	"fmt"
	"strings"

	"github.com/fleeper2133/tasks-app/internal/domain"
	"github.com/jmoiron/sqlx"
)

type TasksPostgres struct {
	db *sqlx.DB
}

func NewTasksPostgres(db *sqlx.DB) *TasksPostgres {
	return &TasksPostgres{db: db}
}

func (r *TasksPostgres) GetAll(userId int) ([]domain.Task, error) {
	var tasks []domain.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", tasksTable)
	err := r.db.Select(&tasks, query, userId)
	return tasks, err
}

func (r *TasksPostgres) Create(task domain.TaskInput, userId int) (int, error) {
	var idTask int
	query := fmt.Sprintf("INSERT INTO %s (title, description, user_id, date_time) VALUES ($1, $2, $3, $4) RETURNING id", tasksTable)
	row := r.db.QueryRow(query, task.Title, task.Description, userId, task.DateTime)
	if err := row.Scan(&idTask); err != nil {
		return 0, err
	}
	return idTask, nil
}

func (r *TasksPostgres) GetById(taskId int, userId int) (domain.Task, error) {
	var task domain.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 and user_id=$2", tasksTable)
	if err := r.db.Get(&task, query, taskId, userId); err != nil {
		return domain.Task{}, err
	}
	return task, nil
}

func (r *TasksPostgres) Delete(taskId int, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 and user_id=$2", tasksTable)
	_, err := r.db.Exec(query, taskId, userId)
	return err
}

func (r *TasksPostgres) Update(taskId int, input domain.TaskUpdate, userId int) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.IsFinish != nil {
		setValues = append(setValues, fmt.Sprintf("is_finish=$%d", argId))
		args = append(args, *input.IsFinish)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id=$%d", tasksTable, setQuery, argId)
	args = append(args, userId)
	_, err := r.db.Exec(query, args...)
	return err

}
