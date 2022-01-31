package routes

import (
	"goer/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Docs
// @version         1.0

// @host      localhost:3000
// @BasePath  /

// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
func MapSwagRoutes(r *gin.Engine) {
	// swagger info
	docs.SwaggerInfo.Title = "API Docs"
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "localhost:3000"
	// docs.SwaggerInfo.BasePath = "/"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// use ginSwagger middleware to serve the API docs
	// DefaultModelsExpandDepth: set -1 to hide models below
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
}
