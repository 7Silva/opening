package handler

import (
	"net/http"

	"github.com/7Silva/openings/functions"
	"github.com/gin-gonic/gin"
)

// @BasePath /v1

// @Summary Get all openings
// @Description Get all openings
// @Tags openings
// @Accept json
// @Produce json
// @Success 200 {object} GetOpeningsResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func GetOpeningsHandler(ctx *gin.Context) {
	openings, err := functions.GetOpenings()
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error getting openings from database")
		return
	}

	if len(openings) == 0 {
		sendError(ctx, http.StatusNotFound, "No openings found")
		return
	}

	sendSuccess(ctx, "getOpenings", openings)
}
