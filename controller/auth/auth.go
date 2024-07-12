package auth

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	jsonformat "github.com/loyalsfc/investrite/jsonFormat"
	"github.com/loyalsfc/investrite/models"
	"github.com/loyalsfc/investrite/response"
	"github.com/loyalsfc/investrite/utils"
)

type AuthHandler struct {
	UserService models.UserService
}

type FormData struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (h AuthHandler) NewUser(ctx *gin.Context) {
	var form FormData

	if ctx.ShouldBind(&form) != nil {
		response.Error(ctx, 403, "invalid form parameter")
		return
	}

	if userExist := h.UserService.IsUserExist(form.Email); userExist {
		response.Error(ctx, 403, "user with the email already exist")
		return
	}

	password, err := utils.HashPassword(form.Password)
	if err != nil {
		response.Error(ctx, 501, "internal error occured")
		return
	}

	user := models.User{
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Email:     form.Email,
		Password:  password,
		UserID:    uuid.New(),
	}

	result := h.UserService.DB.Create(&user)

	if result.Error != nil {
		response.Error(ctx, 401, "An error occured")
		return
	}

	response.Success(ctx, "user added successfully", user)
}

func (h AuthHandler) Signin(ctx *gin.Context) {
	var user FormData
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
