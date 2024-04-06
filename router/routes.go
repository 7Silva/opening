package router

import (
	"github.com/7Silva/openings/handler"
	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {
	handler.InitHandler()

	v1 := r.Group("/v1")
	{
		openings := v1.Group("/openings")
		{
			openings.POST("", handler.CreateOpeningHandler)
			openings.GET("", handler.GetOpeningsHandler)
			openings.GET("/:id", handler.GetOpeningHandler)
			openings.PUT("/:id", handler.UpdateOpeningHandler)
			openings.DELETE("/:id", handler.DeleteOpeningHandler)
		}
	}
}
