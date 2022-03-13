package registerQuery

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/registerQuery"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("/", h.Create)
	r.GET("/", h.GetAll)
	r.GET("/execute/:id", h.Execute)
	r.GET("/:id", h.Get)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)

}
