package registerPropertyToUser

import (
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var registerProperty models.RegisterPropertyToUser
	err := c.Bind(&registerProperty)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	userId, err := models.GetUserID(c)
	registerProperty.UserID = *userId
	err = h.service.Create(&registerProperty)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, registerProperty)
}

func (h *Handler) Delete(c *gin.Context) {
	var registerProperty models.RegisterPropertyToUser
	propertyID, err := uuid.Parse(c.Param("id"))
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	registerProperty.RegisterPropertyID = propertyID
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	userId, err := models.GetUserID(c)
	registerProperty.UserID = *userId
	err = h.service.Delete(&registerProperty)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, nil)
}
