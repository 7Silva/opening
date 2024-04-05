package router

import "github.com/gin-gonic/gin"

func Start() {
	r := gin.Default()

	// initRoutes(r)

	r.Run(":8080")
}
