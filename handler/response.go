package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, gin.H{
		"message":   message,
		"errorCode": statusCode,
	})
}
func sendSuccessResponse(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation %s was successful", operation),
		"data":    data,
	})
}
