package representativesdomains

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/patientsdomains"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("/to-domain", h.AddToDomain)
}
