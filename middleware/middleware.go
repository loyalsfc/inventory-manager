package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/loyalsfc/investrite/response"
	"github.com/loyalsfc/investrite/utils"
	"gorm.io/gorm"
)

type Middleware struct {
	DB *gorm.DB
}

type handlerFunc func(*gin.Context, uuid.UUID)

func (m *Middleware) MiddlewareAuth(handler handlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := utils.GetAccessToken(&ctx.Request.Header)

		if err != nil {
			response.Error(ctx, 301, fmt.Sprintf("error %v", err))
			return
		}

		tokens, err := utils.ParseToken(token)
		if err != nil {
			fmt.Println(err)
			response.Error(ctx, 301, fmt.Sprintf("error %v", err))
			return
		}

		userId, err := uuid.Parse(fmt.Sprintf("%v", tokens))

		if err != nil {
			response.Error(ctx, 301, fmt.Sprintf("error %v", err))
			return
		}

		handler(ctx, userId)
	}
}
