package middleware

import (
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
