package handler

import (
	"fmt"
	"net/http"

	"github.com/diegoliveiraa/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Delete a opening
// @Description Delete a opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpening(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendErrorResponse(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}
	if err := db.Where("id = ?", id).First(&opening).Error; err != nil {
		logger.Errorf("error finding opening: %v", err.Error())
		sendErrorResponse(ctx, http.StatusNotFound, fmt.Sprintf("opening not found: %v", err.Error()))
		return
	}
	if err := db.Delete(&opening).Error; err != nil {
		logger.Errorf("error deleting opening: %v", err.Error())
		sendErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening: %v", err.Error()))
		return
	}
	sendSuccessResponse(ctx, "delete opening", opening)
	logger.Infof("opening deleted: %v", opening)
}
