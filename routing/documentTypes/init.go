package documentTypes

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/documentTypes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB) {
	var h = handler.CreateHandler(db)
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
