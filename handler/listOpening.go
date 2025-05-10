package handler

import (
	"net/http"

	"github.com/diegoliveiraa/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary List all opening
// @Description List all openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
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
