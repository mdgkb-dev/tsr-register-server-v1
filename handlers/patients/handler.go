package patients

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/helpers/writers/errorshelper"
	"mdgkb/tsr-tegister-server-v1/helpers/writers/validators"
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Patient
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	// err = item.FillModelInfoCreate(c, h.helper.Token)

	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	err = h.helper.DB.WithinTransaction(c, func(ctx context.Context) error {
		return h.service.Create(c.Request.Context(), &item)
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

func (h *Handler) FTSP(c *gin.Context) {
	data, err := h.service.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, models.FTSPAnswer{Data: data, FTSP: *h.helper.SQL.ExtractFTSP(c.Request.Context())})
}

func (h *Handler) GetAll(c *gin.Context) {
	items, err := h.service.GetAll(c.Request.Context())

	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := h.service.Get(c.Request.Context(), id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetBySnilsNumber(c *gin.Context) {
	snils := c.Param("snils")
	err := validators.SnilsCheckControlSum(snils)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorshelper.NewHttpError(err.Error(), http.StatusBadRequest, err.Error()))
		return
	}
	item, existsInDomain, err := h.service.GetBySnilsNumber(c.Request.Context(), snils)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, struct {
		Item           *models.Patient `json:"item"`
		ExistsInDomain bool            `json:"existsInDomain"`
	}{Item: item, ExistsInDomain: existsInDomain})
}

func (h *Handler) GetActualAnthropomethry(c *gin.Context) {
	id := c.Param("id")
	height, weight, date, err := S.GetActualAnthropomethry(c.Request.Context(), id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, struct {
		Weight uint       `json:"weight"`
		Height uint       `json:"height"`
		Date   *time.Time `json:"date"`
	}{Weight: weight, Height: height, Date: date})
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
