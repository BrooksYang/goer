package commonRequest

import (
	"mime/multipart"

	"github.com/goer-project/goer/form"
)

type UploadRequest struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

func (req UploadRequest) Messages() form.ValidatorMessages {
	return form.ValidatorMessages{
		"Image.required": "Image is required",
	}
}
