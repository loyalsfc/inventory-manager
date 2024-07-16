package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/loyalsfc/investrite/data"
	"github.com/loyalsfc/investrite/models"
	"github.com/loyalsfc/investrite/response"
	"github.com/loyalsfc/investrite/utils"
)

type UserHandler struct {
	UserService models.UserService
}

func (u UserHandler) GetAllUsers(ctx *gin.Context, userId uuid.UUID) {
	users, err := u.UserService.GetUsers()

	if err != nil {
		response.Error(ctx, 401, fmt.Sprintf("err %v", err))
		return
	}

	response.Success(ctx, "users retried successfully", users)
}

func (u UserHandler) UpdateRole(ctx *gin.Context, userId uuid.UUID) {
	var data data.UpdateUserParams

	ctx.Bind(&data)

	userRole, err := u.UserService.GetUserRole(userId)

	if err != nil {
		response.Error(ctx, 404, fmt.Sprintf("%v", err))
		return
	}

	if userRole != utils.AdminRole {
		response.PermissionError(ctx)
		return
	}

	if err := u.UserService.UpdateUserRole(data.Email, data.Role); err != nil {
		response.Error(ctx, 401, fmt.Sprintf("%v", err))
		return
	}

	response.Success(ctx, "user role updated successfully", nil)
}

func (u UserHandler) DeleteUser(ctx *gin.Context, userId uuid.UUID) {
	id, err := utils.GetIDInRoute(ctx, "userID")

	if err != nil {
		response.Error(ctx, 400, "bad request")
		return
	}

	userRole, err := u.UserService.GetUserRole(userId)

	if err != nil {
		response.Error(ctx, 404, fmt.Sprintf("%v", err))
		return
	}

	if id != userId && userRole != utils.AdminRole {
		response.PermissionError(ctx)
		return
	}

	if err := u.UserService.DeleteUser(id); err != nil {
		response.Error(ctx, 404, fmt.Sprintf("%v", err))
		return
	}

	response.Success(ctx, "user deleted successfully", nil)
}
