package handler

import (
	"fmt"
	"net/http"

	"github.com/7Silva/openings/schemas"
	"github.com/gin-gonic/gin"
)

func sendResponse(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.JSON(code, gin.H{
		"message": message,
		"data":    data,
	})
}

func sendError(ctx *gin.Context, code int, message string) {
	sendResponse(ctx, code, message, nil)
}

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	message := fmt.Sprintf("Operation from handler: %s successful", operation)
	sendResponse(ctx, http.StatusOK, message, data)
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

type DeleteOpeningResponse struct {
	Data    string `json:"data"`
	Message string `json:"message"`
}

type OpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type GetOpeningsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}
