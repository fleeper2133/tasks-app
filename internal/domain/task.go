package domain

import (
	"errors"
)

type Task struct {
	Id              int    `json:"id" db:"id"`
	Title           string `json:"title" db:"title" binding:"required"`
	Description     string `json:"description" db:"description"`
	UserId          int    `db:"user_id" json:"user_id"`
	IsFinish        bool   `db:"is_finish" json:"is_finish" binding:"required"`
	DateTime        string `db:"date_time" json:"date_time" format:"date-time"`
	DateTimeCreated string `db:"date_time_created" json:"date_time_created" format:"date-time"`
}

type TaskInput struct {
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	DateTime    string `db:"date_time" json:"date_time" binding:"required" format:"date-time"`
}

type TaskUpdate struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	IsFinish    *bool   `json:"is_finish"`
}

func (tu TaskUpdate) Validate() error {
	if tu.Title == nil && tu.Description == nil && tu.IsFinish == nil {
		return errors.New("update structures has no values")
	}
	return nil
}
