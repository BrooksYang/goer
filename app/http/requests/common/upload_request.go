package commonRequest

import (
	"mime/multipart"

	"goer/pkg/http"
)

type UploadRequest struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

func (req UploadRequest) Messages() http.ValidatorMessages {
	return http.ValidatorMessages{
		"Image.required": "Image is required",
	}
}
