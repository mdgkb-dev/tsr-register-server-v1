package fileInfo

import (
	"mdgkb/tsr-tegister-server-v1/helpers/uploadHelper"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"

	handler "mdgkb/tsr-tegister-server-v1/handlers/fileInfo"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB, uploader uploadHelper.Uploader) {
	h := handler.CreateHandler(db, &uploader)
	r.GET("/:id", h.Download)
}
