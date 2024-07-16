package response

import "github.com/gin-gonic/gin"

func Success(ctx *gin.Context, message string, payload interface{}) {
	ctx.JSON(200, gin.H{
		"status":  "success",
		"message": message,
		"payload": payload,
	})
}

func Error(ctx *gin.Context, statuscode int, message string) {
	ctx.JSON(statuscode, gin.H{
		"status":  "unsuccessful",
		"message": message,
		"payload": nil,
	})
}

func PermissionError(ctx *gin.Context) {
	ctx.JSON(403, gin.H{
		"status":  "unsuccessful",
		"message": "permission not granted",
		"payload": nil,
	})
}
