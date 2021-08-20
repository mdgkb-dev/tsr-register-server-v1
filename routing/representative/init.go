package representative

import (
"github.com/gin-gonic/gin"
"github.com/uptrace/bun"
handler "mdgkb/tsr-tegister-server-v1/handlers/representative"
"mdgkb/tsr-tegister-server-v1/helpers"
_ "github.com/go-pg/pg/v10/orm"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB) {
var h = handler.CreateHandler(db)
r.GET("/", h.GetAll)
r.GET("/:id", h.Get)
r.POST("/", h.Create)
r.DELETE("/:id", h.Delete)
r.PUT("/:id", h.Update)
}
