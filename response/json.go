package response

import "github.com/gin-gonic/gin"

func Success(ctx *gin.Context, message string, payload interface{}) {
	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": message,
		"payload": payload,
	})
}

func Error(ctx *gin.Context, status int, message string) {
	ctx.JSON(403, gin.H{
		"status":  "unsuccessful",
		"message": message,
		"payload": nil,
	})
}
