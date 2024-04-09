package researches

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/researches"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h *handler.Handler) {
	r.GET("", h.GetAll)
	r.GET("/xlsx/:research-id/:patient-id", h.Xlsx)
	r.GET("/:id", h.Get)
	r.POST("/ftsp", h.FTSP)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
}
