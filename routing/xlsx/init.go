package xlsx

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/xlsx"
	"mdgkb/tsr-tegister-server-v1/helpers/xlsxHelper"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB, xlsxHelper xlsxHelper.IXlsxHelper) {
	var h = handler.CreateHandler(db, xlsxHelper)
	r.GET("/register-query/:id", h.RegisterQuery)
}
