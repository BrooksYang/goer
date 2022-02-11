package middleware

import (
	"time"

	"goer/global"
	"goer/global/errno"
	"goer/pkg/helpers"
	"goer/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

// CheckOpenRequest 签名校验
// 1. 时间戳校验
// 2. 随机数校验
// 3. Api Key 校验（是否有效，ip是否有效）
// 4. 签名校验
func CheckOpenRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 是否开启「签名校验」
		if !global.Config.Open.Enabled {
			c.Next()
			return
		}

		requestBody, _ := c.GetRawData()
		logFields := []zap.Field{
			zap.String("ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("uri", c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("params", string(requestBody)),
		}

		// 1. 时间戳校验
		timestamp := c.Request.FormValue("timestamp")
		if cast.ToInt64(timestamp) < time.Now().Unix()-30 {
			global.Logger.Open.Warn("timestamp error", logFields...)
			response.Fail(c, errno.IllegalRequest)
			c.Abort()
			return
		}

		// 2. 随机数校验
		nonce := c.Request.FormValue("nonce")
		nonceExists := !helpers.Empty(global.Cache.Get(nonce))
		if nonce == "" || nonceExists {
			global.Logger.Open.Warn("nonce error", logFields...)
			response.Fail(c, errno.IllegalRequest)
			c.Abort()
			return
		}

		// 3. Api Key 校验（是否有效，ip是否有效）
		apiKey := c.Request.FormValue("access_key")
		if apiKey != global.Config.Open.ApiKey {
			global.Logger.Open.Warn("api key error", logFields...)
			response.Fail(c, errno.IllegalRequest)
			c.Abort()
			return
		}

		// check ip
		if helpers.Empty(global.Config.Open.Ip) || helpers.Contains([]string{"", "*"}, global.Config.Open.Ip[0]) {
			c.Next()
			return
		}

		if !helpers.Contains(global.Config.Open.Ip, c.ClientIP()) {
			global.Logger.Open.Warn("ip error", logFields...)
			response.Fail(c, errno.IllegalRequest)
			c.Abort()
			return
		}

		c.Next()
	}
}
