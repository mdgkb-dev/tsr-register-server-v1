package commissionsdoctors

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/commissionsdoctors"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("", h.Create)
	r.GET("", h.GetAll)
	r.GET("/:id", h.Get)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
}
