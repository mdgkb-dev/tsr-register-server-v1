package middleware

import (
	"github.com/gin-gonic/gin"
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Extract the access token metadata
		_, err := models.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
	}
}
