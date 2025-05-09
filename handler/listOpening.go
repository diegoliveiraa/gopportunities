package handler

import (
	"net/http"

	"github.com/diegoliveiraa/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpenings(ctx *gin.Context) {
	openings := []schemas.Opening{}
	if err := db.Find(&openings).Error; err != nil {
		logger.Errorf("error finding openings: %v", err.Error())
		sendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccessResponse(ctx, "list openings", openings)
	logger.Infof("openings found: %v", openings)
}
