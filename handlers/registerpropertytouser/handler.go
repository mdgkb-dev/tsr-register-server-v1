package registerpropertytouser

import (
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var registerProperty models.RegisterPropertyToUser
	err := c.Bind(&registerProperty)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	userID, err := h.helper.Token.GetUserID(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	registerProperty.UserID = *userID
	err = h.service.Create(&registerProperty)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, registerProperty)
}

func (h *Handler) Delete(c *gin.Context) {
	var registerProperty models.RegisterPropertyToUser
	propertyID, err := uuid.Parse(c.Param("id"))
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	registerProperty.RegisterPropertyID = propertyID
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	userID, err := h.helper.Token.GetUserID(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	registerProperty.UserID = *userID
	err = h.service.Delete(&registerProperty)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, nil)
}
