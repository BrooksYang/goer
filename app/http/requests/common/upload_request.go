package commonRequest

import (
	"mime/multipart"

	"goer/pkg/form"
)

type UploadRequest struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

func (req UploadRequest) Messages() form.ValidatorMessages {
	return form.ValidatorMessages{
		"Image.required": "Image is required",
	}
}
