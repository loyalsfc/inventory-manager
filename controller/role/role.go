package role

import (
	"github.com/google/uuid"
	"github.com/loyalsfc/investrite/database"
	"github.com/loyalsfc/investrite/models"
	"github.com/loyalsfc/investrite/utils"
)

func GetUserRole(userId uuid.UUID) (utils.UserRole, error) {
	db, _ := database.InitDB()

	userService := models.UserService{
		DB: db,
	}

	userRole, err := userService.GetUserRole(userId)

	if err != nil {
		return "", err
	}

	return userRole, nil
}
