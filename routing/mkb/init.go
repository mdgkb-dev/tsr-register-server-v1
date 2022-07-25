package mkb

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/mkb"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAllClasses)
	r.PUT("/:id/", h.Update)
	r.GET("/groups/:classId", h.GetGroupByClassID)
	r.GET("/sub-groups/:groupId", h.GetGroupChildrens)
	r.GET("/sub-sub-groups/:subGroupId", h.GetSubGroupChildrens)
	r.GET("/groups", h.GetGroupsBySearch)
	r.GET("/diagnosis", h.GetDiagnosisBySearch)
	r.GET("/sub-diagnosis", h.GetSubDiagnosesBySearch)
	r.GET("/diagnosis/byGroupId/:groupId", h.GetDiagnosisByGroupID)
	r.GET("/diagnosis/:diagnosisId", h.GetSubDiagnosisByDiagnosisID)
	r.GET("/concrete-diagnosis/:diagnosisId", h.GetConcreteDiagnosisBySubDiagnosisID)
	r.GET("/concrete-diagnosis", h.GetConcreteDiagnosisBySearch)
}
