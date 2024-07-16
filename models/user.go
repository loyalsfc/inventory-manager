package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/loyalsfc/investrite/data"
	"github.com/loyalsfc/investrite/utils"
	"gorm.io/gorm"
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
	Role      string    `json:"role" gorm:"column:role"`
	UserID    uuid.UUID `json:"user_id" gorm:"column:user_id;unique;not null"`
}

type APIUser struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	UserID    uuid.UUID `json:"user_id"`
}

func (u UserService) CreateUser(form data.FormData) (*User, error) {
	if userExist := u.IsUserExist(form.Email); userExist {
		return nil, errors.New("user with the email already exist")
	}

	password, err := utils.HashPassword(form.Password)
	if err != nil {
		return nil, err
	}

	user := User{
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Email:     form.Email,
		Password:  password,
		UserID:    uuid.New(),
	}

	result := u.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u User) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Model(u).Update("role", utils.ViewerRole)
	return
}

func (u UserService) GetUsers() ([]APIUser, error) {
	var users []APIUser

	if result := u.DB.Model(&User{}).Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (u UserService) IsUserExist(email string) bool {
	result := u.DB.Where("email = ?", email).First(&User{})

	if result.Error == nil {
		return true
	} else {
		return false
	}
}

func (u UserService) GetUser(email string) (*User, error) {
	var user User
	result := u.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u UserService) GetUserById(userId uuid.UUID) (*User, error) {
	var user User
	result := u.DB.Where("user_id = ?", userId).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u UserService) UpdateUserRole(email string, newRole utils.UserRole) error {
	_, changedUserErr := u.GetUser(email)

	if changedUserErr != nil {
		return changedUserErr
	}

	if isValidRole := utils.IsValidRole(newRole); !isValidRole {
		return errors.New("the provided role is invalid")
	}

	if result := u.DB.Model(&User{}).Where("email = ?", email).Update("role", newRole); result.Error != nil {
		return result.Error
	}

	return nil
}

func (u UserService) GetUserRole(userId uuid.UUID) (utils.UserRole, error) {
	user, err := u.GetUserById(userId)

	if err != nil {
		return "", err
	}

	return utils.UserRole(user.Role), nil
}

func (u UserService) DeleteUser(userId uuid.UUID) error {
	if result := u.DB.Where("user_id = ?", userId).Delete(&User{}); result.Error != nil {
		return result.Error
	}

	return nil
}
