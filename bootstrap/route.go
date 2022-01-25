package bootstrap

import (
	"goer/app/http/middleware"
	"goer/routes"

	"github.com/gin-gonic/gin"
)

// Setup router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Global middlewares
	registerGlobalMiddleWare(r)

	// gin-swagger
	routes.MapSwagRoutes(r)

	return r
}

// Register global middlewares
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middleware.Cors(),
	)
}
