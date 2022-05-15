package patients

import (
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Patient
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = item.FillModelInfoCreate(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.Create(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.historyService.Create(&item, models.RequestTypeInsert)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	err := h.service.setQueryFilter(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusUnauthorized) {
		return
	}
	withDisabilities := c.Query("withDisabilities")
	if withDisabilities != "" {
		items, err := h.service.GetDisabilities()
		if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
			return
		}
		c.JSON(http.StatusOK, items)
		return
	}
	query := c.Query("query")
	if query != "" {
		items, err := h.service.GetBySearch(&query)
		if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
			return
		}
		c.JSON(http.StatusOK, items)
		return
	}
	onlyNames := c.Query("only-names")
	if onlyNames != "" {
		items, err := h.service.GetOnlyNames()
		if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
			return
		}
		c.JSON(http.StatusOK, items)
		return
	}
	items, err := h.service.GetAll()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(&id, true)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(&id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	item, err := h.service.Get(&id, true)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.historyService.Create(item, models.RequestTypeDelete)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Patient
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = item.FillModelInfoUpdate(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.Update(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.historyService.Create(&item, models.RequestTypeUpdate)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAllHistory(c *gin.Context) {
	id := c.Param("id")
	items, err := h.historyService.GetAll(&id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetHistory(c *gin.Context) {
	id := c.Param("id")
	items, err := h.historyService.Get(&id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}
