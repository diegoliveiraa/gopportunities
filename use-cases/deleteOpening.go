package usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete opening",
	})
}
