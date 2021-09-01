package meta

import (
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCount(c *gin.Context) {
	table := c.Param("table")
	items, err := h.service.GetCount(&table)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetSchema(c *gin.Context) {
	c.JSON(http.StatusOK, h.service.GetSchema())
}
