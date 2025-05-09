package router

import (
	"github.com/diegoliveiraa/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(rounter *gin.Engine) {
	v1 := rounter.Group("/api/v1")
	{
		// Show opening
		v1.GET("/opening", handler.ShowopeningHandler)
		// Create opening
		v1.POST("/opening", handler.CreateopeningHandler)
		// Delete opening
		v1.DELETE("/opening", handler.DeleteopeningHandler)
		// Update opening
		v1.PUT("/opening", handler.UpdateopeningHandler)
		// Show all openings
		v1.GET("/openings", handler.ShowallopeningHandler)
	}
}
