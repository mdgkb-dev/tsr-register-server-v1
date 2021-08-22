package uploadHelper

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type Uploader interface {
	GetUploaderPath() *string
	Upload(*gin.Context, []*multipart.FileHeader, string) error
}

type LocalUploader struct {
	UploadPath *string
}

func NewLocalUploader(path *string) *LocalUploader {
	return &LocalUploader{
		UploadPath: path,
	}
}

func (u *LocalUploader) Upload(c *gin.Context, file []*multipart.FileHeader, name string) (err error) {
	uploadPath := u.GetUploaderPath()
	path := filepath.Join(*uploadPath, name)
	parts := strings.Split(path, string(os.PathSeparator))
	err = os.MkdirAll(filepath.Join(parts[:len(parts)-1]...), os.ModePerm)
	if err != nil {
		return err
	}
	err = c.SaveUploadedFile(file[0], path)
	if err != nil {
		return err
	}
	return nil
}

func (u *LocalUploader) GetUploaderPath() *string {
	return u.UploadPath
}
