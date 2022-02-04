package response

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Ts   int64       `json:"ts"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS = 0
)

func Result(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Ts:   time.Now().UnixMilli(),
		Data: data,
	})
}

func ErrnoResult(c *gin.Context, errno *Errno, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: errno.Code,
		Msg:  errno.Msg,
		Ts:   time.Now().UnixMilli(),
		Data: data,
	})
}

func Success(c *gin.Context) {
	Result(c, SUCCESS, map[string]interface{}{}, "Success")
}

func Msg(c *gin.Context, message string) {
	Result(c, SUCCESS, map[string]interface{}{}, message)
}

func Data(c *gin.Context, data interface{}) {
	Result(c, SUCCESS, data, "OK")
}

func Detail(c *gin.Context, data interface{}, message string) {
	Result(c, SUCCESS, data, message)
}

func Fail(c *gin.Context, errno *Errno) {
	ErrnoResult(c, errno, map[string]interface{}{})
}

func FailWithMsg(c *gin.Context, code int, message string) {
	Result(c, code, map[string]interface{}{}, message)
}

func FailWithData(c *gin.Context, errno *Errno, data interface{}) {
	ErrnoResult(c, errno, data)
}

// BadRequest 400
func BadRequest(c *gin.Context, err error, msg ...string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": defaultMessage("Bad Request", msg...),
		"error":   err.Error(),
	})
}

// Unauthorized 401
func Unauthorized(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": defaultMessage("Unauthorized", msg...),
	})
}

// ValidationError 422
func ValidationError(c *gin.Context, errors map[string][]string) {
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
		"message": "Validation error",
		"errors":  errors,
	})
}

// defaultMessage
func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMsg
	}
	return
}
