package handler

import (
	usecases "github.com/diegoliveiraa/gopportunities/use-cases"
	"github.com/gin-gonic/gin"
)

func CreateopeningHandler(ctx *gin.Context) {
	usecases.CreateOpening(ctx)
}

func UpdateopeningHandler(ctx *gin.Context) {
	usecases.UpdateOpening(ctx)
}

func DeleteopeningHandler(ctx *gin.Context) {
	usecases.DeleteOpening(ctx)
}
func ShowopeningHandler(ctx *gin.Context) {
	usecases.ShowOpening(ctx)
}

func ShowallopeningHandler(ctx *gin.Context) {
	usecases.ShowOpening(ctx)
}
