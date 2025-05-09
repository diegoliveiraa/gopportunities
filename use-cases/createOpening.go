package usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create opening",
	})
}
