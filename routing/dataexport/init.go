package dataexport

import (
	"github.com/gin-gonic/gin"

	handler "mdgkb/tsr-tegister-server-v1/handlers/dataexport"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("", h.Export)
}
