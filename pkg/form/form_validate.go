package form

import (
	"encoding/json"
	"errors"

	"goapp/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
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

func ParseErrors(request FormRequest, err error) map[string][]string {
	data := map[string][]string{}

	for _, v := range err.(validator.ValidationErrors) {
		errField := strcase.ToSnake(v.Field())
		errMsg := v.Error()

		if message, exist := request.Messages()[v.Field()+"."+v.Tag()]; exist {
			errMsg = message
		}

		if _, ok := data[errField]; ok {
			data[errField] = append(data[errField], errMsg)
		} else {
			data[errField] = []string{errMsg}
		}
	}

	return data
}

func Validate(c *gin.Context, request FormRequest) bool {
	if err := c.ShouldBind(request); err != nil {
		// Unmarshal error
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			response.BadRequest(c, errors.New("illegal parameter"))
			return false
		}

		// Validation error
		response.ValidationError(c, ParseErrors(request, err))
		return false
	}

	return true
}
