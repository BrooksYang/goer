package common

import (
	v1 "goapp/app/http/controllers/v1"
	"goapp/pkg/response"

	"github.com/gin-gonic/gin"
)

type CommonController struct {
	v1.BaseController
}

// Ping
// @BasePath  /
// PingExample godoc
// @Summary   ping server
// @Schemes
// @Description  do ping
// @Tags         Common
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  "Ping"
// @Router       /v1/common/ping [get]
func (common *CommonController) Ping(c *gin.Context) {
	response.Msg(c, "pong")
}
