package fileinfos

import (
	"mdgkb/tsr-tegister-server-v1/models"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/uploadHelper"
)

func (s *FilesService) GetFullPath(fileSystemPath *string) *string {
	return s.helper.Uploader.GetFullPath(fileSystemPath)
}

func (s *FilesService) Upload(c *gin.Context, item *models.FileInfo, files map[string][]*multipart.FileHeader) (err error) {
	for i, file := range files {
		if i == item.ID.UUID.String() {
			item.FileSystemPath = uploadHelper.BuildPath(&i)
			err = s.helper.Uploader.Upload(c, file, &item.FileSystemPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
