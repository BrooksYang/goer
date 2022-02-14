package routes

import (
	"goer/app/http/controllers/v1/common"
	"goer/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func MapCommonRoutes(r *gin.Engine) {

	// v1 group
	v1 := r.Group("/v1/common")

	// controllers
	commonController := new(common.CommonController)

	// Ping test
	v1.GET("/ping", commonController.Ping)

	/**
	|--------------------------------------------------------------------------
	| File
	|--------------------------------------------------------------------------
	|
	| Here is where you can register API routes for your application. These
	| routes are loaded by the RouteServiceProvider within a group which
	| is assigned the "api" middleware group. Enjoy building your API!
	|
	*/

	// Controllers
	fileController := common.FileController{}

	// Upload
	v1.POST("upload", middleware.Auth(), fileController.Upload)
}
