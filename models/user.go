package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	AdminRole      = "admin"
	SupervisorRole = "supervisor"
	OperatorRole   = "operator"
	ViewerRole     = "viewer"
)

type UserService struct {
	DB *gorm.DB
}

type User struct {
	gorm.Model
	FirstName string    `json:"first_name" gorm:"column:first_name;not null"`
	LastName  string    `json:"last_name" gorm:"column:last_name;not null"`
	Email     string    `json:"email" gorm:"column:email;unique;not null"`
	Password  string    `json:"password" gorm:"column:password;not nul"`
	Role      string    `json:"role" gorm:"column:role;default:viewer;not null"`
	UserID    uuid.UUID `json:"user_id" gorm:"column:user_id;unique;not null"`
}

func (u UserService) IsUserExist(email string) bool {
	result := u.DB.Where("email = ?", email).First(&User{})

	if result.Error == nil {
		return true
	} else {
		return false
	}
}

func (u UserService) GetUser(email string) (User, error) {
	var user User
	result := u.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
