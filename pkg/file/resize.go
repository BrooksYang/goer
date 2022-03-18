package file

import (
	"mime/multipart"
	"os"
	"strings"

	"github.com/disintegration/imaging"
)

func Resize(dir string, filename string, file *multipart.FileHeader) (string, error) {
	// Get path
	dir = strings.TrimRight(dir, "/") + "/"
	path := dir + filename

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
	resizedPath := dir + randomNameFromUploadFile(file)
	err = imaging.Save(src, resizedPath)
	if err != nil {
		return "", err
	}

	// Remove old file
	err = os.Remove(path)
	if err != nil {
		return "", err
	}

	return resizedPath, nil
}
