package patientsresearches

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/patientsresearches"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("", h.GetAll)
	r.GET("/:id", h.Get)
	r.GET("patient-research/:patientId/:researchId", h.GetPatientResearch)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
}
