package commissionsdrugapplications

import (
	"mdgkb/tsr-tegister-server-v1/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

func (s *FilesService) Upload(c *gin.Context, item *models.CommissionDrugApplication, files map[string][]*multipart.FileHeader) (err error) {
	//for i, file := range files {
	//	err := s.helper.Uploader.Upload(c, file, item.SetFilePath(&i))
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}
