package auth

import (
	handler "mdgkb/tsr-tegister-server-v1/handlers/auth"

	"github.com/gin-gonic/gin"
)

// Init func
func Init(r *gin.RouterGroup, h handler.IHandler) {
	r.POST("/login", h.Login)
	r.POST("/register", h.Register)
	r.POST("/refresh-token", h.RefreshToken)
}
