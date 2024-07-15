package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/loyalsfc/investrite/controller/auth"
	"github.com/loyalsfc/investrite/controller/categories"
	"github.com/loyalsfc/investrite/controller/items"
	"github.com/loyalsfc/investrite/controller/orders"
	"github.com/loyalsfc/investrite/middleware"
	"github.com/loyalsfc/investrite/models"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": uuid.New(),
		})
	})

	middlware := &middleware.Middleware{
		DB: db,
	}

	userService := &models.UserService{
		DB: db,
	}

	authHandler := &auth.AuthHandler{
		UserService: *userService,
	}

	authRoutes := r.Group("/user")
	authRoutes.POST("/new", authHandler.NewUser)
	authRoutes.POST("/signin", authHandler.Signin)

	categoryModel := &models.CategoryModel{
		DB: db,
	}

	categoryHandler := &categories.CategoryHandler{
		CategoryService: *categoryModel,
	}

	categoryRoute := r.Group("/category")
	categoryRoute.POST("/new-category", middlware.MiddlewareAuth(categoryHandler.NewCategory))
	categoryRoute.PUT("/:id", categoryHandler.EditCategory)
	categoryRoute.DELETE("/:id", categoryHandler.DeleteCategory)
	categoryRoute.GET("/", categoryHandler.GetCategories)

	productService := &models.ProductService{
		DB: db,
	}

	productHandler := &items.ProductHandler{
		ProductService: *productService,
	}

	productRoute := r.Group("/product")
	productRoute.POST("/new-product", middlware.MiddlewareAuth(productHandler.NewProduct))
	productRoute.GET("/", middlware.MiddlewareAuth(productHandler.GetProducts))
	productRoute.GET("/:productID", middlware.MiddlewareAuth(productHandler.GetProduct))
	productRoute.PUT("/:productID", middlware.MiddlewareAuth(productHandler.UpdateProduct))
	productRoute.DELETE("/:productID", middlware.MiddlewareAuth(productHandler.DeleteProduct))
	productRoute.GET("/increase-quantity/:productID", middlware.MiddlewareAuth(productHandler.IncreaseProductQuantity))
	productRoute.GET("/decrease-quantity/:productID", middlware.MiddlewareAuth(productHandler.DecreaseProductQuantity))

	orderService := models.OrderService{
		DB: db,
	}

	orderHandler := orders.OrderHandler{
		OrderService: orderService,
	}

	orderRoutes := r.Group("/order")
	orderRoutes.POST("/new", middlware.MiddlewareAuth(orderHandler.NewOrder))
	orderRoutes.GET("/", middlware.MiddlewareAuth(orderHandler.GetOrders))
	orderRoutes.GET("/:orderId", middlware.MiddlewareAuth(orderHandler.GetOrder))
	orderRoutes.DELETE("/:orderId", middlware.MiddlewareAuth(orderHandler.DeleteOrder))

	return r
}
