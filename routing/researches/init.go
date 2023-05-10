package researches

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/researches"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("", h.GetAll)
	r.GET("/xlsx", h.GetXlsx)
	r.GET("/:id", h.Get)
	r.POST("/", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)

	r.GET("/value-types", h.GetValueTypes)
}