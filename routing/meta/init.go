package meta

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"

	handler "mdgkb/tsr-tegister-server-v1/handlers/meta"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB) {
	var h = handler.CreateHandler(db)
	r.GET("/count/:table", h.GetCount)
	r.GET("/schema", h.GetSchema)
}
