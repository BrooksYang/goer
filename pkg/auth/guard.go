package auth

import (
	"goer/app/models/user"

	"github.com/gin-gonic/gin"
)

// Id Get user id
func Id(c *gin.Context) (id uint64) {
	return c.GetUint64("auth.user_id")
}

// User Get user info
func User(c *gin.Context) user.User {
	authUser, ok := c.MustGet("auth.user").(user.User)
	if !ok {
		return user.User{}
	}

	return authUser
}
