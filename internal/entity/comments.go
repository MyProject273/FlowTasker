package entity

type Comment struct {
	BaseEntity
	TaskID uint `json:"task_id" gorm:"not null"`
	Task   Task `gorm:"foreignKey:TaskID" json:"-"`
	UserID uint `json:"user_id" gorm:"not null"`
	User   User `gorm:"foreignKey:UserID" json:"user"`

	Content string `json:"content" gorm:"not null;type:varchar(255)"`
}
