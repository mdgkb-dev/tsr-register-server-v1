package auth

import (
	"mdgkb/tsr-tegister-server-v1/handlers/users"
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	var user *models.UserAccount
	err := c.Bind(&user)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	item, err := h.service.Register(user)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Login(c *gin.Context) {
	var item models.UserAccount
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	//err = h.validator.Login(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	res, err := h.service.Login(&item, false)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) LoginAs(c *gin.Context) {
	var item models.UserAccount
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	res, err := h.service.Login(&item, true)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(c *gin.Context) {
	//_, err := h.helper.Token.ExtractTokenMetadata(c.Request)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, "unauthorized")
	//	return
	//}
	//delErr := helpers.DeleteTokens(metadata, h.redis)
	//if delErr != nil {
	//	c.JSON(http.StatusUnauthorized, delErr.Error())
	//	return
	//}
	c.JSON(http.StatusOK, "Successfully logged out")
}

type refreshToken struct {
	RefreshToken string `json:"refreshToken"`
	UserID       string `json:"userId"`
}

func (h *Handler) RefreshToken(c *gin.Context) {
	t := refreshToken{}
	err := c.Bind(&t)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	//userId, err := h.helper.Token.ExtractTokenMetadata(c.Request, "user_id")
	//if h.helper.HTTP.HandleError(c, err) {
	//	return
	//}
	user, err := users.CreateService(h.helper).Get(t.UserID)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	tokens, err := h.helper.Token.RefreshToken(t.RefreshToken, user)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, tokens)
}
