package fileInfo

import (
	"github.com/gin-gonic/gin"
	handler "mdgkb/tsr-tegister-server-v1/handlers/fileInfo"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/:id", h.Download)
}
