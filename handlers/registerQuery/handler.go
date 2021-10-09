package registerQuery

import (
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var query models.RegisterQuery
	err := c.Bind(&query)

	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	err = h.service.Create(&query)

	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, query)
}

func (h *Handler) GetAll(c *gin.Context) {
	queries, err := h.service.GetAll()

	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, queries)
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	query, err := h.service.Get(&id)

	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, query)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(&id)

	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var query models.RegisterQuery
	err := c.Bind(&query)

	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	err = h.service.Update(&query)

	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, query)
}
