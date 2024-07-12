package categories

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/loyalsfc/investrite/models"
	"github.com/loyalsfc/investrite/response"
	"github.com/loyalsfc/investrite/utils"
)

type CategoryHandler struct {
	CategoryService models.CategoryModel
}

type CategoryStruct struct {
	Name string `json:"name"`
}

func (c CategoryHandler) NewCategory(ctx *gin.Context, userId uuid.UUID) {
	var name CategoryStruct

	ctx.Bind(&name)

	category, err := c.CategoryService.CreateCategory(name.Name)

	if err != nil {
		response.Error(ctx, 403, err.Error())
		return
	}

	response.Success(ctx, "new category created", category)
}

func (c CategoryHandler) EditCategory(ctx *gin.Context) {
	catID, err := utils.GetIdFromParams(ctx)

	if err != nil {
		response.Error(ctx, 403, err.Error())
		return
	}

	var category CategoryStruct
	ctx.Bind(&category)

	editErr := c.CategoryService.EditCategory(category.Name, catID)

	if editErr != nil {
		response.Error(ctx, 401, editErr.Error())
		return
	}

	response.Success(ctx, "category edit successful", nil)
}

func (c CategoryHandler) DeleteCategory(ctx *gin.Context) {
	catID, err := utils.GetIdFromParams(ctx)

	if err != nil {
		response.Error(ctx, 403, err.Error())
		return
	}

	deleteErr := c.CategoryService.DeleteCategory(catID)

	if deleteErr != nil {
		response.Error(ctx, 301, deleteErr.Error())
		return
	}

	response.Success(ctx, "delete successful", nil)
}

func (c CategoryHandler) GetCategories(ctx *gin.Context) {
	categories, err := c.CategoryService.CategoryList()

	if err != nil {
		response.Error(ctx, 401, fmt.Sprintf("%v", err))
		return
	}

	response.Success(ctx, "category list succesful", categories)
}
