package drugdozes

import (
	"github.com/gin-gonic/gin"

	handler "mdgkb/tsr-tegister-server-v1/handlers/drugdozes"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/", h.Create)
	r.POST("/calculate-needing", h.CalculateNeeding)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
}
