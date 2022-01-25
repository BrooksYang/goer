package auth

import (
	"strings"

	"goer/app/models/user"
	"goer/global"

	"github.com/gin-gonic/gin"
)

// Get Authorization Bearer Token
func GetToken(c *gin.Context) string {
	bearToken := c.Request.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

// Get user id
func Id(c *gin.Context) (id uint64) {
	token := GetToken(c)
	if token == "" {
		return
	}

	// Parse token
	j := NewJWT()
	claims, _ := j.ParseToken(token)
	if claims == nil {
		return
	}

	return claims.ID
}

// Get user info
func User(c *gin.Context) (authUser user.User) {
	userId := Id(c)
	if userId == 0 {
		return
	}

	global.DB.First(&authUser, userId)
	return
}
