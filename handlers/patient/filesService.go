package patient

import (
	"github.com/gin-gonic/gin"
	"mdgkb/tsr-tegister-server-v1/models"
	"mime/multipart"
)

func (s *FilesService) Upload(c *gin.Context, item *models.Patient, files map[string][]*multipart.FileHeader) (err error) {
	for i, file := range files {
		newPath := item.SetFilePath(&i)
		err := s.uploader.Upload(c, file, *newPath)
		if err != nil {
			return err
		}
	}
	return nil
}
