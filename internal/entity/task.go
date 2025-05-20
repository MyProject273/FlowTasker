package entity

import "time"

type TaskStatus string

const (
	TaskStatusPending  TaskStatus = "pending"
	TaskStatusProgress TaskStatus = "in_progress"
	TaskStatusDone     TaskStatus = "done"
)

type Task struct {
	BaseEntity
	Title       string     `json:"title" gorm:"type:varchar(255);not null"`
	Description string     `json:"description" gorm:"type:varchar(255)"`
	Status      TaskStatus `json:"status" gorm:"type:varchar(20);not null;default:'pending'"`
	Priority    uint       `json:"priority" gorm:"check:priority BETWEEN 1 AND 5;not null"`
	DueDate     time.Time  `json:"due_date"`

	//Ref to project
	ProjectID uint    `json:"project_id" `
	Project   Project `gorm:"foreignKey:ProjectID" json:"-"`

	TaskAssign []TaskAssign `json:"task_assigns" gorm:"foreignKey:TaskID"`
	Comment    []Comment    `json:"comments" gorm:"foreignKey:TaskID"`
}
