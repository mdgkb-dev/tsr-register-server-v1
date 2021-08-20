package mkb

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
	handler "mdgkb/tsr-tegister-server-v1/handlers/mkb"
)

// Init func
func Init(r *gin.RouterGroup, db *bun.DB) {
	var h = handler.CreateHandler(db)
	r.GET("/", h.GetAllClasses)
	r.PUT("/:id/", h.Update)
	r.GET("/groups/:classId", h.GetGroupByClassId)
	r.GET("/sub-groups/:groupId", h.GetGroupChildrens)
	r.GET("/sub-sub-groups/:subGroupId", h.GetSubGroupChildrens)
	r.GET("/groups", h.GetGroupsBySearch)
	r.GET("/diagnosis", h.GetDiagnosisBySearch)
	r.GET("/diagnosis/byGroupId/:groupId", h.GetDiagnosisByGroupId)
	r.GET("/diagnosis/:diagnosisId", h.GetSubDiagnosisByDiagnosisId)
}
