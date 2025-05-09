package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(rounter *gin.Engine) {
	v1 := rounter.Group("/api/v1")
	{
		// Show opening
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Oppening",
			})
		})
		// Create opening
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{
				"message": "POST Oppenings",
			})
		})
		// Delete opening
		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "DELETE Oppenings",
			})
		})
		// Update opening
		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "PUT Oppenings",
			})
		})
		// Show all openings
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "List all Openings",
			})
		})
	}
}
