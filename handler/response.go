package handler

import (
	"fmt"
	"net/http"

	"github.com/diegoliveiraa/gopportunities/schemas"
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

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}
type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
type ShowOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
