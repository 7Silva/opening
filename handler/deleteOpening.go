package handler

import (
	"fmt"
	"net/http"

	"github.com/7Silva/openings/functions"
	"github.com/gin-gonic/gin"
)

// @BasePath /v1

// @Summary Delete an opening
// @Description Delete a new job an opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id path string true "Opening ID"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings/{id} [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}

	if !functions.OpeningExists(id) {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Opening with id: %s not found", id))
		return
	}

	if !functions.DeleteOpening(id) {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error deleting opening with id: %s on database", id))
		return
	}

	sendSuccess(ctx, "deleteOpening", fmt.Sprintf("Opening with id: %s deleted", id))
}
