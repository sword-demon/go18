package task

import "context"

type Service interface {
	Run(context.Context) (*Task, error)
}

type Task struct {
	Id string `json:"id" gorm:"column:id;type:varchar(60)" description:"id"`
}

type TaskSpec struct {
}
