package middleware

import (
	"goer/global/errno"
	"goer/pkg/auth"
	"goer/pkg/response"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get auth user
		userInfo := auth.User(c)
		if userInfo.ID == 0 {
			response.Fail(errno.InvalidToken, c)
			c.Abort()
			return
		}

		// account invalid
		if !userInfo.IsValid {
			response.Fail(errno.AccountLocked, c)
			c.Abort()
			return
		}

		c.Next()
	}
}
