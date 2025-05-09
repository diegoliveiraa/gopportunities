package router

import (
	"github.com/diegoliveiraa/gopportunities/docs"
	"github.com/diegoliveiraa/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
