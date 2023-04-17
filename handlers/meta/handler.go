package meta

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pro-assistance/pro-assister/projecthelper"
)

func (h *Handler) GetCount(c *gin.Context) {
	table := c.Param("table")
	items, err := h.service.GetCount(&table)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetSchema(c *gin.Context) {
	c.JSON(http.StatusOK, projecthelper.SchemasLib)
}
