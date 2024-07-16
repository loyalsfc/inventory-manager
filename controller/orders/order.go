package orders

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/loyalsfc/investrite/controller/role"
	"github.com/loyalsfc/investrite/data"
	"github.com/loyalsfc/investrite/models"
	"github.com/loyalsfc/investrite/response"
	"github.com/loyalsfc/investrite/utils"
)

type OrderHandler struct {
	OrderService models.OrderService
}

func (o OrderHandler) NewOrder(ctx *gin.Context, userId uuid.UUID) {
	var params data.OrderParams
	ctx.Bind(&params)

	userRole, err := role.GetUserRole(userId)

	if err != nil {
		response.Error(ctx, 403, fmt.Sprintf("%v", err))
		return
	}

	if utils.RoleLevel(userRole) < 2 {
		response.PermissionError(ctx)
		return
	}

	order, err := o.OrderService.CreateOrder(params)

	if err != nil {
		response.Error(ctx, 403, err.Error())
		return
	}

	response.Success(ctx, "order created successfully", order)
}

func (o OrderHandler) GetOrders(ctx *gin.Context, userId uuid.UUID) {
	userRole, err := role.GetUserRole(userId)

	if err != nil {
		response.Error(ctx, 403, fmt.Sprintf("%v", err))
		return
	}

	if utils.RoleLevel(userRole) < 2 {
		response.PermissionError(ctx)
		return
	}

	orders, err := o.OrderService.GetAllOrders()

	if err != nil {
		response.Error(ctx, 400, err.Error())
		return
	}

	response.Success(ctx, "orders retrieved successfully", orders)
}

func (o OrderHandler) GetOrder(ctx *gin.Context, userId uuid.UUID) {
	userRole, err := role.GetUserRole(userId)

	if err != nil {
		response.Error(ctx, 403, fmt.Sprintf("%v", err))
		return
	}

	if utils.RoleLevel(userRole) < 2 {
		response.PermissionError(ctx)
		return
	}

	id, err := utils.GetIDInRoute(ctx, "orderId")

	if err != nil {
		response.Error(ctx, 401, err.Error())
		return
	}
	order, err := o.OrderService.FindOrder(id)

	if err != nil {
		response.Error(ctx, 401, err.Error())
		return
	}

	response.Success(ctx, "order retrieved successfully", order)
}

func (o OrderHandler) DeleteOrder(ctx *gin.Context, userId uuid.UUID) {
	userRole, err := role.GetUserRole(userId)

	if err != nil {
		response.Error(ctx, 403, fmt.Sprintf("%v", err))
		return
	}

	if utils.RoleLevel(userRole) < 4 {
		response.PermissionError(ctx)
		return
	}

	id, err := utils.GetIDInRoute(ctx, "orderId")

	if err != nil {
		response.Error(ctx, 401, err.Error())
		return
	}
	deleteErr := o.OrderService.DeleteOrder(id)

	if deleteErr != nil {
		response.Error(ctx, 401, deleteErr.Error())
		return
	}

	response.Success(ctx, "order deleted successfully", nil)
}
