package entity

type User struct {
	BaseEntity
	Name           string `json:"name" gorm:"type:varchar(255)"`
	Email          string `json:"email" gorm:"type:varchar(255);not null;unique;index"`
	PasswordHashed string `json:"password_hashed" gorm:"type:varchar(255);not null"`
	Role           string `json:"role" gorm:"type:varchar(255);not null"`
}
