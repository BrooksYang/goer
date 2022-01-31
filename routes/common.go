package routes

import (
	"goer/app/http/controllers/v1/common"

	"github.com/gin-gonic/gin"
)

// Map demo routes
func MapCommonRoutes(r *gin.Engine) {

	// v1 group
	v1 := r.Group("/v1/common")

	// controllers
	commonController := new(common.CommonController)

	// Ping test
	v1.GET("/ping", commonController.Ping)
}
