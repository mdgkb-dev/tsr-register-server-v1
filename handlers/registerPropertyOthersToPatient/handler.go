package registerPropertyOthersToPatient

import (
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.RegisterPropertyOthersToPatient
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

//
func (h *Handler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
