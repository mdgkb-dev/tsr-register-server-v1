package xlsx

import (
	"github.com/gin-gonic/gin"

	handler "mdgkb/tsr-tegister-server-v1/handlers/xlsx"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.GET("/register-query/:id", h.RegisterQuery)
}
