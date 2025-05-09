package handler

import (
	"net/http"

	"github.com/diegoliveiraa/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpening(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	logger.Infof("request received: %v", request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		sendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccessResponse(ctx, "create opening", opening)
	logger.Infof("opening created: %v", opening)
}
