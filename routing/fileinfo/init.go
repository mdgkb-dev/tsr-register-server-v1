package fileinfo

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/fileinfo"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/:id", h.Download)
}