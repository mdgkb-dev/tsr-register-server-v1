package patient

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func (s *FilesService) Upload(c *gin.Context, item *models.Patient, files map[string][]*multipart.FileHeader) (err error) {
	for i, file := range files {
		newPath := item.SetFilePath(&i)
		fmt.Println(newPath)
		err := s.uploader.Upload(c, file, *newPath)
		if err != nil {
			return err
		}
	}
	return nil
}
