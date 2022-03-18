package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/goer-project/goer-utils/helpers"
)

func Put(data []byte, to string) error {
	err := ioutil.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}

	return true
}

func GetFileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func SaveUploadedFile(c *gin.Context, file *multipart.FileHeader) (string, error) {
	// Mkdir
	storagePath := "storage/public"
	dirName := fmt.Sprintf("/uploads/%s/", time.Now().Format("2006/01/02"))
	_ = os.MkdirAll(storagePath+dirName, 0755)

	// Random filename
	fileName := randomNameFromUploadFile(file)

	path := storagePath + dirName + fileName
	if err := c.SaveUploadedFile(file, path); err != nil {
		return "", err
	}

	// Open image
	src, err := imaging.Open(path, imaging.AutoOrientation(true))
	if err != nil {
		return "", err
	}

	// Resize ratio
	resizeRatio := getResizeRatio(file)
	width := float64(src.Bounds().Size().X) * resizeRatio

	// Resize
	src = imaging.Resize(src, int(width), 0, imaging.Lanczos)
	log.Println(src.Bounds().Size())
	resizedFilename := randomNameFromUploadFile(file)
	resizedPath := storagePath + dirName + resizedFilename
	err = imaging.Save(src, resizedPath)
	if err != nil {
		return "", err
	}

	// Remove old file
	err = os.Remove(path)
	if err != nil {
		return "", err
	}

	resizedPath = strings.Replace(resizedPath, "public/", "", 1)

	return resizedPath, nil
}

func randomNameFromUploadFile(file *multipart.FileHeader) string {
	return helpers.RandomString(40) + filepath.Ext(file.Filename)
}

func getResizeRatio(file *multipart.FileHeader) float64 {
	// < 100k
	if file.Size < 1024*100 {
		return 1
	}

	// 100k - 300k
	if file.Size <= 1024*300 {
		return 0.8
	}

	// 300k - 500k
	if file.Size <= 1024*500 {
		return 0.6
	}

	// 500k - 1M
	if file.Size <= 1024*1024 {
		return 0.5
	}

	// 1M - 5M
	if file.Size <= 1024*1024*5 {
		return 0.3
	}

	// > 5M
	return 0.1
}
