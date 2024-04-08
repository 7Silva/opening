package router

import (
	docs "github.com/7Silva/openings/docs"
	"github.com/7Silva/openings/functions"
	"github.com/7Silva/openings/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRoutes(r *gin.Engine) {
	functions.InitDB()
	basePath := "/v1"

	docs.SwaggerInfo.Title = "Openings API"
	docs.SwaggerInfo.BasePath = basePath

	v1 := r.Group(basePath)
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
