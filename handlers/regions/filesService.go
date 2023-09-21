package regions

import (
	"mdgkb/tsr-tegister-server-v1/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func (s *FilesService) Upload(_ *gin.Context, _ *models.Region, _ map[string][]*multipart.FileHeader) (err error) {
	//for i, file := range files {
	//	err := s.helper.Uploader.Upload(c, file, item.SetFilePath(&i))
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}
