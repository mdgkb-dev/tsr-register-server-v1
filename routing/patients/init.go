package patients

import (
	"github.com/gin-gonic/gin"

	handler "mdgkb/tsr-tegister-server-v1/handlers/patients"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("", h.GetAll)
	r.GET("/actual-anthropometry/:id", h.GetActualAnthropomethry)
	r.GET("/:id", h.Get)
	r.GET("/by-snils/:snils", h.GetBySnilsNumber)
	r.POST("", h.Create)
	r.DELETE("/:id", h.Delete)
	r.PUT("/:id", h.Update)
}
