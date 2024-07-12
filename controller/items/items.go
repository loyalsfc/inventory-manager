package items

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/loyalsfc/investrite/data"
	"github.com/loyalsfc/investrite/models"
	"github.com/loyalsfc/investrite/response"
	"github.com/loyalsfc/investrite/utils"
)

type ProductHandler struct {
	ProductService models.ProductService
}

func (h ProductHandler) NewProduct(ctx *gin.Context, userId uuid.UUID) {
	var params data.AddProductParams

	ctx.Bind(&params)

	product, err := h.ProductService.CreateProduct(&params)

	if err != nil {
		response.Error(ctx, 403, fmt.Sprintf("%v", err))
		return
	}

	response.Success(ctx, "product added successfully", product)
}

func (h ProductHandler) GetProduct(ctx *gin.Context, userId uuid.UUID) {
	productId, err := utils.GetIDInRoute(ctx, "productID")

	if err != nil {
		response.Error(ctx, 403, fmt.Sprintf("%v", err))
		return
	}

	product, err := h.ProductService.GetProductById(productId)

	if err != nil {
		response.Error(ctx, 404, fmt.Sprintf("%v", err))
		return
	}

	response.Success(ctx, "product retrieved successfully", product)
}

func (h ProductHandler) GetProducts(ctx *gin.Context, userId uuid.UUID) {
	products, err := h.ProductService.GetAllProducts()

	if err != nil {
		response.Error(ctx, 404, fmt.Sprintf("%v", err))
		return
	}

	response.Success(ctx, "product retrieved successfully", products)
}

func (h ProductHandler) UpdateProduct(ctx *gin.Context, userId uuid.UUID) {
	productId, err := utils.GetIDInRoute(ctx, "productID")

	var params data.AddProductParams

	ctx.Bind(&params)

	if err != nil {
		response.Error(ctx, 403, fmt.Sprintf("%v", err))
		return
	}

	productErr := h.ProductService.UpdateProduct(productId, &params)

	if productErr != nil {
		response.Error(ctx, 404, fmt.Sprintf("%v", productErr))
		return
	}

	response.Success(ctx, "product updated successfully", nil)
}

func (h ProductHandler) DeleteProduct(ctx *gin.Context, userId uuid.UUID) {
	productId, err := utils.GetIDInRoute(ctx, "productID")

	if err != nil {
		response.Error(ctx, 403, fmt.Sprintf("%v", err))
		return
	}

	productErr := h.ProductService.DeleteProduct(productId)

	if productErr != nil {
		response.Error(ctx, 404, fmt.Sprintf("%v", productErr))
		return
	}

	response.Success(ctx, "product deleted", nil)
}

func (h ProductHandler) IncreaseProductQuantity(ctx *gin.Context, userId uuid.UUID) {
	productId, err := utils.GetIDInRoute(ctx, "productID")

	if err != nil {
		response.Error(ctx, 403, fmt.Sprintf("%v", err))
		return
	}

	qty, productErr := h.ProductService.IncrementQuantity(productId)

	if productErr != nil {
		response.Error(ctx, 404, fmt.Sprintf("%v", productErr))
		return
	}

	response.Success(ctx, "product quantity increased", qty)
}

func (h ProductHandler) DecreaseProductQuantity(ctx *gin.Context, userId uuid.UUID) {
	productId, err := utils.GetIDInRoute(ctx, "productID")

	if err != nil {
		response.Error(ctx, 403, fmt.Sprintf("%v", err))
		return
	}

	qty, productErr := h.ProductService.DecreaseQuantity(productId)

	if productErr != nil {
		response.Error(ctx, 404, fmt.Sprintf("%v", productErr))
		return
	}

	response.Success(ctx, "product quantity decreased", qty)
}
