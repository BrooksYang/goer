package bootstrap

import (
	"goer/app/http/middleware"
	"goer/global"
	"goer/routes"

	"github.com/gin-gonic/gin"
)

// Setup router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Global middlewares
	registerGlobalMiddleWare(r)

	// Static file server
	r.StaticFS("storage", gin.Dir("./storage/public", false))

	// gin-swagger
	if global.Config.Swag.ApiDoc {
		routes.MapSwagRoutes(r)
	}

	// Common routes
	routes.MapCommonRoutes(r)

	// api
	routes.MapApiRoutes(r)

	return r
}

// Register global middlewares
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middleware.Cors(),             // 允许跨域请求
		middleware.ForceUserAgent(),   // 强制附带 User-Agent
		middleware.LimitMethod("1-S"), // 限制请求方法，GET 不限制，其它每秒1次
		middleware.LimitIP("10-S"),    // 限制 IP 请求，每秒10次
		middleware.LogRequest(),       // 请求日志
	)
}
