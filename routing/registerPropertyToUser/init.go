package registerPropertyToUser

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/registerPropertyToUser"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB) {
	var h = handler.CreateHandler(db)
	r.POST("/", h.Create)
	r.DELETE("/:id", h.Delete)
}
