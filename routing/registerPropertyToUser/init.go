package registerPropertyToUser

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	handler "mdgkb/tsr-tegister-server-v1/handlers/registerPropertyToUser"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("/", h.Create)
	r.DELETE("/:id", h.Delete)
}
