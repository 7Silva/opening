package handler

import (
	"net/http"

	"github.com/7Silva/openings/functions"
	"github.com/7Silva/openings/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /v1

// @Summary Create an opening
// @Description Create an opening
// @Tags openings
// @Accept json
// @Produce json
// @Param body body CreateOpeningRequest true "Request body"
// @Success 200 {object} OpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
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

	opening, err := functions.CreateOpening(opening)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "createOpening", opening)
}
