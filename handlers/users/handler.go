package users

import (
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(c *gin.Context) {
	err := h.service.setQueryFilter(c)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	items, err := h.service.GetAll()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	item, err := h.service.Get(c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Update(c *gin.Context) {
	var item models.User
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.Update(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Create(c *gin.Context) {
	var item models.User
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.Create(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

type FavouriteForm struct {
	ID string `json:"id"`
}
