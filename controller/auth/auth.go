package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/loyalsfc/investrite/data"
	jsonformat "github.com/loyalsfc/investrite/jsonFormat"
	"github.com/loyalsfc/investrite/models"
	"github.com/loyalsfc/investrite/response"
	"github.com/loyalsfc/investrite/utils"
)

type AuthHandler struct {
	UserService models.UserService
}

func (h AuthHandler) NewUser(ctx *gin.Context) {
	var form data.FormData

	if ctx.ShouldBind(&form) != nil {
		response.Error(ctx, 403, "invalid form parameter")
		return
	}

	user, err := h.UserService.CreateUser(form)

	if err != nil {
		response.Error(ctx, 401, fmt.Sprintf("%v", err))
		return
	}

	response.Success(ctx, "user added successfully", user)
}

func (h AuthHandler) Signin(ctx *gin.Context) {
	var user data.FormData
	ctx.Bind(&user)

	if userExist := h.UserService.IsUserExist(user.Email); !userExist {
		response.Error(ctx, 404, "invalid email or password")
		return
	}

	userInfo, _ := h.UserService.GetUser(user.Email)

	if isValid := utils.ComparePassword(user.Password, userInfo.Password); !isValid {
		response.Error(ctx, 403, "invalid password")
		return
	}

	signString, err := utils.GenerateToken(userInfo.UserID)

	if err != nil {
		response.Error(ctx, 500, fmt.Sprintf("error %v", err))
		return
	}

	response.Success(ctx, "login successful", jsonformat.SignInToSignIn(userInfo, signString))
}
