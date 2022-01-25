package http

import (
	"goer/global/errno"
	"goer/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FormRequest interface {
	Messages() ValidatorMessages
}

type ValidatorMessages map[string]string

func ParseError(request FormRequest, err error) string {
	for _, v := range err.(validator.ValidationErrors) {
		if message, exist := request.Messages()[v.Field()+"."+v.Tag()]; exist {
			return message
		}

		return v.Error()
	}

	return "Parameter error"
}

func ParseErrors(request FormRequest, err error) map[string]string {
	data := map[string]string{}

	for _, v := range err.(validator.ValidationErrors) {
		if message, exist := request.Messages()[v.Field()+"."+v.Tag()]; exist {
			data[v.Field()] = message
			continue
		}

		data[v.Field()] = v.Error()
	}

	return data
}

func Validate(c *gin.Context, request FormRequest) bool {
	if err := c.ShouldBind(request); err != nil {
		response.FailWithMsg(c, errno.ValidationError.Code, ParseError(request, err))
		return false
	}

	return true
}
