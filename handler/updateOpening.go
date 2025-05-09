package handler

import (
	"net/http"

	"github.com/diegoliveiraa/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpening(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)
	logger.Infof("request received: %v", request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		sendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendErrorResponse(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.Where("id = ?", id).First(&opening).Error; err != nil {
		logger.Errorf("error finding opening: %v", err.Error())
		sendErrorResponse(ctx, http.StatusNotFound, "opening not found!")
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error updating opening: %v", err.Error())
		sendErrorResponse(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendSuccessResponse(ctx, "update opening", opening)
	logger.Infof("opening updated: %v", opening)
}
