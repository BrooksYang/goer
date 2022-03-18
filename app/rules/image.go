package rules

import (
	"mime/multipart"

	"goer/global/errno"
	"goer/pkg/form"
	"goer/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/goer-project/goer-utils/file"
)

// Image struct for image information and storage location
type Image struct {
	Mime string `validate:"required,oneof=image/png image/jpg image/jpeg"`
	Size int32  `validate:"required,gt=0,lte=5242880"`
}

func (req Image) Messages() form.ValidatorMessages {
	return form.ValidatorMessages{
		"Mime.required": "Image is required",
		"Mime.oneof":    "Image only support png, jpg, jpeg",
		"Size.lte":      "The image must not be greater than 5MB",
	}
}

func ValidateImage(c *gin.Context, header *multipart.FileHeader) bool {

	fileContent, err := header.Open()
	mime, err := file.GetContentType(fileContent)
	if err != nil {
		response.BadRequest(c, err)
		return false
	}

	img := Image{
		Mime: mime,
		Size: int32(header.Size),
	}

	validate := validator.New()
	err = validate.Struct(img)
	if err == nil {
		return true
	}

	response.FailWithMsg(c, errno.ValidationError.Code, form.ParseError(img, err))

	return false
}
