package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null;type:varchar(100);index:unique"`
	Email    string `gorm:"not null;type:varchar(100);index:unique"`
	Password string `gorm:"not null;type:varchar(100)"`
	Age      uint   `gorm:"not null;"`
	// Role     uint   `gorm:"not null;"`
}

type Todos struct {
	gorm.Model
	Status      uint   `gorm:"not null"`
	Description string `gorm:"not null"`
}

type Photo struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Caption  string `gorm:"not null"`
	PhotoUrl string `gorm:"not null"`

	// FK users
	UserID uint
	User   User
}

type Comment struct {
	gorm.Model

	// FK users
	UserID uint
	User   User

	// FK photos
	PhotoID uint
	Photo   Photo

	Message string
}

type SocialMedia struct {
	gorm.Model
	Name           string
	SocialMediaUrl string
	// FK users
	UserID uint
	User   User
}
