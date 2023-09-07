package questions

import (
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Question
	err := c.Bind(&item)

	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	err = h.service.Create(c, &item)

	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	//ctx, err := models.User{}.InjectClaims(c.Request, h.helper.Token)
	//if h.helper.HTTP.HandleError(c, err) {
	//	return
	//}

	//fq, err := h.helper.SQL.CreateQueryFilter(c)
	//if h.helper.HTTP.HandleError(c, err) {
	//	return
	//}
	//h.helper.SQL.InjectQueryFilter(ctx, fq)
	items, err := h.service.GetAll(c.Request.Context())

	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(c, id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c, id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Question
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.Update(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}
