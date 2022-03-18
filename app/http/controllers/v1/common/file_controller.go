package common

import (
	v1 "goer/app/http/controllers/v1"
	commonRequest "goer/app/http/requests/common"
	"goer/app/rules"
	"goer/global/errno"
	"goer/pkg/form"
	"goer/pkg/response"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	v1.BaseController
}

// Upload
// @Summary   Upload image
// @Security  Bearer
// @Tags      Common
// @Accept    multipart/form-data
// @Produce   json
// @Param     image  formData  file    true  "image"
// @Success   200    {string}  string  "OK"
// @Router    /v1/common/upload [POST]
func (*FileController) Upload(c *gin.Context) {
	var request commonRequest.UploadRequest
	if ok := form.Validate(c, &request); !ok {
		return
	}
	if ok := rules.ValidateImage(c, request.Image); !ok {
		return
	}

	path, err := form.SaveUploadedFile(c, request.Image)
	if err != nil {
		response.Fail(c, errno.InternalServerError)
	}

	response.Data(c, path)
}
