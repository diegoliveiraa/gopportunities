package router

import (
	"github.com/diegoliveiraa/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(rounter *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	v1 := rounter.Group("/api/v1")
	{
		// Show opening
		v1.GET("/opening", handler.ShowOpening)
		// Create opening
		v1.POST("/opening", handler.CreateOpening)
		// Delete opening
		v1.DELETE("/opening", handler.DeleteOpening)
		// Update opening
		v1.PUT("/opening", handler.UpdateOpening)
		// Show all openings
		v1.GET("/openings", handler.ListOpenings)
	}
}
