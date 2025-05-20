package entity

import "time"

type TaskAssign struct {
	TaskID uint `json:"task_id" gorm:"not null"`
	UserID uint `json:"user_id" gorm:"not null"`

	Task       Task      `json:"-" gorm:"foreignKey:TaskID"`
	User       User      `json:"user" gorm:"foreignKey:UserID"`
	AssignedAt time.Time `json:"assigned_at"`
}
