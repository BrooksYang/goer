package middleware

import (
	"errors"

	"goer/pkg/response"

	"github.com/gin-gonic/gin"
)

func ForceUserAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Request.Header["User-Agent"]) == 0 {
			response.BadRequest(c, errors.New("User-Agent not found"))
			return
		}

		c.Next()
	}
}
