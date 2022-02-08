package rules

import (
	"mime/multipart"

	"goer/global/errno"
	"goer/pkg/http"
	"goer/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Image struct for image information and storage location
type Image struct {
	Mime string `validate:"required,oneof=image/png image/jpg image/jpeg"`
	Size int32  `validate:"required,gt=0,lte=5242880"`
}

func (req Image) Messages() http.ValidatorMessages {
	return http.ValidatorMessages{
		"Mime.required": "Image is required",
		"Mime.oneof":    "Image only support png, jpg, jpeg",
		"Size.lte":      "The image must not be greater than 5MB",
	}
}

func ValidateImage(c *gin.Context, header *multipart.FileHeader) bool {

	img := Image{
		Mime: header.Header.Get("Content-Type"),
		Size: int32(header.Size),
	}

	validate := validator.New()
	err := validate.Struct(img)
	if err == nil {
		return true
	}

	response.FailWithMsg(c, errno.ValidationError.Code, http.ParseError(img, err))

	return false
}
