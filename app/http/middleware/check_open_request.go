package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	"goer/global"
	"goer/global/errno"
	"goer/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/goer-project/goer-utils/helpers"
	"github.com/goer-project/goer-utils/xhttp"
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

		// Get body
		requestBody, _ := c.GetRawData()
		var body map[string]interface{}
		err := json.Unmarshal(requestBody, &body)
		if err != nil {
			response.Fail(c, errno.IllegalRequest)
			c.Abort()
			return
		}

		// Reset request body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		logFields := []zap.Field{
			zap.String("ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("uri", c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("params", string(requestBody)),
		}

		// 1. 时间戳校验
		timestamp := body["timestamp"]
		if cast.ToInt64(timestamp) < time.Now().Unix()-global.Config.Open.TTL {
			global.Logger.Open.Warn("timestamp error", logFields...)
			response.Fail(c, errno.IllegalRequest)
			c.Abort()
			return
		}

		// 2. 随机数校验
		nonce := cast.ToString(body["nonce"])
		nonceExists := global.Cache.Has(nonce)
		if nonce == "" || nonceExists {
			global.Logger.Open.Warn("nonce error", logFields...)
			response.Fail(c, errno.IllegalRequest)
			c.Abort()
			return
		}

		// Cache nonce
		global.Cache.Set(nonce, true, time.Second*time.Duration(global.Config.Open.TTL))

		// 3. Api Key 校验（是否有效，ip是否有效）
		apiKey := cast.ToString(body["access_key"])
		if apiKey != global.Config.Open.ApiKey {
			global.Logger.Open.Warn("api key error", logFields...)
			response.Fail(c, errno.IllegalRequest)
			c.Abort()
			return
		}

		// check ip
		if len(global.Config.Open.Ip) > 0 && !helpers.Contains(global.Config.Open.Ip, c.ClientIP()) {
			global.Logger.Open.Warn("ip error", logFields...)
			response.Fail(c, errno.IllegalRequest)
			c.Abort()
			return
		}

		// 4. 签名校验
		sign := cast.ToString(body["sign"])
		delete(body, "sign")
		resign := xhttp.Sign(body, global.Config.Open.ApiSecret)
		if sign == "" || sign != resign {
			logFields = append(logFields, zap.String("resign", resign))
			global.Logger.Open.Warn("sign error", logFields...)
			response.Fail(c, errno.IllegalRequest)
			c.Abort()
			return
		}

		c.Next()
	}
}
