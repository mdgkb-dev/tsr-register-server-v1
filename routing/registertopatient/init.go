package registertopatient

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/registertopatient"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.DELETE("/:id", h.Delete)
}
