package handler

import (
	"net/http"
	"strconv"

	"github.com/7Silva/openings/functions"
	"github.com/7Silva/openings/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /v1

// @Summary Update an opening
// @Description Update an opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id path string true "Opening ID"
// @Param opening body UpdateOpeningRequest true "Opening data to Update"
// @Success 200 {object} OpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings/{id} [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "ID is required")
		return
	}

	var opening schemas.Opening
	idUint, _ := strconv.ParseUint(id, 10, 64)
	if err := ctx.ShouldBindJSON(&opening); err != nil {
		sendError(ctx, http.StatusBadRequest, "Invalid JSON provided")
		return
	}

	opening.ID = uint(idUint)
	updatedOpening, err := functions.UpdateOpening(opening)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "updateOpening", updatedOpening)
}
