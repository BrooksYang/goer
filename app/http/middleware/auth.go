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
			response.Fail(c, errno.InvalidToken)
			c.Abort()
			return
		}

		// account invalid
		if !userInfo.IsValid {
			response.Fail(c, errno.AccountLocked)
			c.Abort()
			return
		}

		c.Next()
	}
}
