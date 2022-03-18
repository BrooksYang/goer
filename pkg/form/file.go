package form

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	filePkg "github.com/goer-project/goer-utils/file"
)

func SaveUploadedFile(c *gin.Context, file *multipart.FileHeader) (string, error) {
	// Mkdir
	storagePath := "storage/public"
	dirName := fmt.Sprintf("/uploads/%s/", time.Now().Format("2006/01/02"))
	_ = os.MkdirAll(storagePath+dirName, 0755)

	// Random filename
	fileName := filePkg.RandomFilename(file)
	dir := storagePath + dirName
	path := dir + fileName
	if err := c.SaveUploadedFile(file, path); err != nil {
		return "", err
	}

	// Open image
	resizedPath, err := filePkg.Resize(dir, fileName, file)
	if err != nil {
		return "", err
	}

	resizedPath = strings.Replace(resizedPath, "public/", "", 1)

	return resizedPath, nil
}
