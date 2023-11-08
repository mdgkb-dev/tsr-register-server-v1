package representatives

import (
	"mdgkb/tsr-tegister-server-v1/helpers/writers/validators"
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Representative
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	//err = item.FillModelInfoCreate(c)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	//err = item.FillModelInfoCreate(c)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.Create(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
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

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(c.Request.Context(), id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Representative
	files, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	// err = item.FillModelInfoUpdate(c, h.helper.Token)
	// if h.helper.HTTP.HandleError(c, err) {
	// 	return
	// }
	err = h.filesService.Upload(c, &item, files)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	err = h.service.Update(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetBySnilsNumber(c *gin.Context) {
	snils := c.Param("snils")
	err := validators.SnilsCheckControlSum(snils)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	h.helper.Validator.Validate()
	item, existsInDomain, err := h.service.GetBySnilsNumber(c.Request.Context(), snils)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, struct {
		Item           *models.Representative `json:"item"`
		ExistsInDomain bool                   `json:"existsInDomain"`
	}{Item: item, ExistsInDomain: existsInDomain})
}
