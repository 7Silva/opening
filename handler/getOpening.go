package handler

import (
	"fmt"
	"net/http"

	"github.com/7Silva/openings/functions"
	"github.com/gin-gonic/gin"
)

// @BasePath /v1

// @Summary Get an opening
// @Description Get an opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id path string true "Opening ID"
// @Success 200 {object} OpeningResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings/{id} [get]
func GetOpeningHandler(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	if !functions.OpeningExists(id) {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("No opening with id: %s found", id))
		return
	}

	opening, err := functions.GetOpening(id)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error getting opening with id: %s from database", id))
		return
	}

	sendSuccess(ctx, "getOpening", opening)
}
