package form

import (
	"encoding/json"
	"errors"

	"goapp/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/goer-project/goer/form"
)

func Validate(c *gin.Context, request form.FormRequest) bool {
	if err := c.ShouldBind(request); err != nil {
		// Unmarshal error
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			response.BadRequest(c, errors.New("illegal parameter"))
			return false
		}

		// Validation error
		response.ValidationError(c, form.ParseErrors(request, err))
		return false
	}

	return true
}
