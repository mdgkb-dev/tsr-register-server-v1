package patients

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Patient
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	//err = item.FillModelInfoCreate(c, h.helper.Token)

	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	err = h.helper.DB.WithinTransaction(c, func(ctx context.Context) error {
		return h.service.Create(c, &item)
	})

	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	//err = h.historyService.Create(&item, models.RequestTypeInsert)
	//if h.helper.HTTP.HandleError(c, err) {
	//	return
	//}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	fq, err := h.helper.SQL.CreateQueryFilter(c)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	//d, err := h.helper.Token.ExtractTokenMetadata(c.Request, "domains_ids")
	//if h.helper.HTTP.HandleError(c, err) {
	//	return
	//}
	ctx := context.WithValue(c, "fq", fq)

	items, err := h.service.GetAll(ctx)
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
	var item models.Patient
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	//err = item.FillModelInfoUpdate(c, h.helper.Token)
	//if h.helper.HTTP.HandleError(c, err) {
	//	return
	//}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.Update(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	//err = h.historyService.Create(&item, models.RequestTypeUpdate)
	//if h.helper.HTTP.HandleError(c, err) {
	//	return
	//}
	c.JSON(http.StatusOK, item)
}
