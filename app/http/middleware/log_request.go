package middleware

import (
	"bytes"
	"net/http"
	"time"

	"goer/global"
	"goer/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/goer-project/goer-utils/helpers"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Response
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		start := time.Now()
		c.Next()

		// User id
		userId := auth.Id(c)

		// Response
		cost := time.Since(start)
		responseStatus := c.Writer.Status()

		// Log content
		logFields := []zap.Field{
			zap.Uint64("user_id", userId),
			zap.String("ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("uri", c.Request.URL.String()),
			zap.Int("status", responseStatus),
			zap.String("time", helpers.MicrosecondsStr(cost)),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("user-agent", c.Request.UserAgent()),
		}

		// Log request params & response except GET
		if c.Request.Method != http.MethodGet {
			requestBody, _ := c.GetRawData()
			logFields = append(logFields, zap.String("params", string(requestBody)))
			logFields = append(logFields, zap.String("response", w.body.String()))
		}

		// Log
		status := cast.ToString(responseStatus)
		if responseStatus >= 500 {
			global.Logger.Request.Error(status, logFields...)
			return
		}

		if responseStatus >= 400 {
			global.Logger.Request.Warn(status, logFields...)
			return
		}

		global.Logger.Request.Info(status, logFields...)
	}
}
