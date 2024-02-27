package auth

import (
	"net/http"

	baseModels "github.com/pro-assistance/pro-assister/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var item *baseModels.UserAccount
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	res, err := S.Register(c.Request.Context(), item.Email, item.Password)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Login(c *gin.Context) {
	var item baseModels.UserAccount
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	res, err := S.Login(c.Request.Context(), item.Email, item.Password)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) RestorePassword(c *gin.Context) {
	var item baseModels.UserAccount
	_, err := h.helper.HTTP.GetForm(c, item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.RestorePassword(c, item.Email)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, err)
}
