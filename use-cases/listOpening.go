package usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOpenings(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all openings",
	})
}
