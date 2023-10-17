package drugdozes

import (
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.DrugDoze
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.Create(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetAll(c *gin.Context) {
	items, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	item, err := S.Get(c.Request.Context(), c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	err := S.Delete(c.Request.Context(), c.Param("id"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.DrugDoze
	_, err := h.helper.HTTP.GetForm(c, &item)

	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.Update(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

type DrugNeedingOptions struct {
	Weight     uint          `json:"weight"`
	Height     uint          `json:"height"`
	Start      *time.Time    `json:"start"`
	End        *time.Time    `json:"end"`
	DrugDozeID uuid.NullUUID `json:"drugDozeId"`
}

func (h *Handler) CalculateNeeding(c *gin.Context) {
	var item DrugNeedingOptions
	_, err := h.helper.HTTP.GetForm(c, &item)

	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	needing, err := S.CalculateNeeding(c, item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, needing)
}
