package registerQuery

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/registerQuery"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func Init(r *gin.RouterGroup, db *bun.DB) {
	var h = handler.CreateHandler(db)
	r.POST("/", h.Create)
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
