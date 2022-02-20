package middleware

import (
	"goer/app/models/user"
	"goer/global"
	"goer/pkg/auth"
	"goer/pkg/response"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		claims, err := auth.NewJWT().ParseToken(c)
		if err != nil {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		if claims.Guard != "" {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		var userInfo user.User
		global.DB.First(&userInfo, claims.ID)

		if userInfo.ID == 0 {
			response.Unauthorized(c)
			c.Abort()
			return
		}

		// Set user info to gin.context
		c.Set("auth.user_id", userInfo.ID)
		c.Set("auth.user", userInfo)

		c.Next()
	}
}
