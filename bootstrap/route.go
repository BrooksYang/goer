package bootstrap

import (
	"goer/app/http/middleware"

	"github.com/gin-gonic/gin"
)

// Setup router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Global middlewares
	registerGlobalMiddleWare(r)

	// Map routes

	return r
}

// Register global middlewares
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middleware.Cors(),
	)
}
