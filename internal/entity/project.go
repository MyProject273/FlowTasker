package entity

type Project struct {
	BaseEntity
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:varchar(255)"`
	OwnerID     uint   `json:"owner_id" gorm:"not null"`
	Owner       User   `json:"owner" gorm:"foreignKey:OwnerID"`

	Task []Task `json:"tasks" gorm:"foreginKey:ProjectID"`
}
