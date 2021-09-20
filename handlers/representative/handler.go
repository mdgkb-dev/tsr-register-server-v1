package representative

import (
	"encoding/json"
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Representative
	form, _ := c.MultipartForm()
	err := json.Unmarshal([]byte(form.Value["form"][0]), &item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.Create(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	query := c.Query("query")
	if query != "" {
		items, err := h.service.GetBySearch(&query)
		if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
			return
		}
		c.JSON(http.StatusOK, items)
		return
	}
	queryFilter, err := httpHelper.CreateQueryFilter(c)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	if queryFilter.Pagination != nil {
		items, err := h.service.GetAll(queryFilter)
		if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
			return
		}
		c.JSON(http.StatusOK, items)
		return
	}
	items, err := h.service.GetOnlyNames()
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(&id)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
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
	var item models.Representative
	files, err := httpHelper.GetForm(c, &item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	err = h.filesService.Upload(c, &item, files)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	err = h.service.Update(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}
