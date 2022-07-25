package meta

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/meta"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/count/:table", h.GetCount)
	r.GET("/schema", h.GetSchema)
}
