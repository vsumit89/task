package models

import "database/sql/driver"

type TaskStatus string

const (
	Pending    TaskStatus = "pending"
	InProgress TaskStatus = "inProgress"
	Completed  TaskStatus = "completed"
)

func (t *TaskStatus) Scan(value interface{}) error {
	*t = TaskStatus(value.([]byte))
	return nil
}

func (e TaskStatus) Value() (driver.Value, error) {
	return string(e), nil
}

type Task struct {
	Base
	Title string	`gorm:"column:title" json:"title"` 
	Description string `gorm:"column:description" json:"description"`
	Status TaskStatus `sql:"status" gorm:"column:status" json:"status"`
}
