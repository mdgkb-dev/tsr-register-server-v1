package mkbConcreteDiagnoses

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/mkbConcreteDiagnoses"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
}
