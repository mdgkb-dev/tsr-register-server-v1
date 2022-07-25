package registerpropertytouser

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/registerpropertytouser"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("/", h.Create)
	r.DELETE("/:id", h.Delete)
}
