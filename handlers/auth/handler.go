package auth

import (
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var user *models.User
	err := c.Bind(&user)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	res, err := h.service.Register(user)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Login(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	res, err := h.service.Login(&user)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Me(c *gin.Context) {
	userID, err := h.helper.Token.GetUserID(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	userStringID := userID.String()
	res, err := h.service.GetUserByID(&userStringID)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(c *gin.Context) {
	//_, err := models.ExtractTokenMetadata(c.Request)
	//if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
	//	return
	//}
	//delErr := helpers.DeleteTokens(metadata, h.redis)
	//if delErr != nil {
	//	c.JSON(http.StatusUnauthorized, delErr.Error())
	//	return
	//}
	c.JSON(http.StatusOK, "Successfully logged out")
}

func (h *Handler) DoesLoginExist(c *gin.Context) {
	login := c.Param("login")
	doesLoginExist, _ := h.service.DoesLoginExist(&login)
	c.JSON(http.StatusOK, &DoesLoginExist{doesLoginExist})
}
func (h *Handler) RefreshToken(c *gin.Context) {
	type refreshToken struct {
		RefreshToken string `json:"refreshToken"`
	}
	t := refreshToken{}
	err := c.Bind(&t)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	tokens, err := h.helper.Token.RefreshToken(t.RefreshToken)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, tokens)
}

// Пока что на проекте нет путей доступа
func (h *Handler) CheckPathPermissions(c *gin.Context) {
	var path string
	err := c.Bind(&path)
	if h.helper.HTTP.HandleError(c, err, http.StatusForbidden) {
		return
	}
	//userRoleId := ""
	if c.Request.Header.Get("token") != "null" {
		_, err := h.helper.Token.GetAccessDetail(c)
		if h.helper.HTTP.HandleError(c, err, http.StatusUnauthorized) {
			return
		}
		//userRoleId = accessDetails.UserRoleID
	}
	//err = h.service.CheckPathPermissions(path, userRoleId)
	//if h.helper.HTTP.HandleError(c, err, http.StatusForbidden) {
	//	return
	//}
	c.JSON(http.StatusOK, nil)
}
