package meta

import (
	"github.com/gin-gonic/gin"
	handler "mdgkb/tsr-tegister-server-v1/handlers/meta"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/count/:table", h.GetCount)
	r.GET("/schema", h.GetSchema)
}
