package pathpermissions

import (
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SavePathPermissions(c *gin.Context) {
	var items models.PathPermissions
	err := c.Bind(&items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.UpsertManyPathPermissions(items)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, err)
}

func (h *Handler) GetAllPathPermissions(c *gin.Context) {
	items, err := h.service.GetAllPathPermissions()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetAllPathPermissionsAdmin(c *gin.Context) {
	err := h.service.setQueryFilter(c)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	items, err := h.service.GetAllPathPermissionsAdmin()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetPathPermissionsByRoleID(c *gin.Context) {
	items, err := h.service.GetPathPermissionsByRoleID(c.Param("roleId"))
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) CheckPathPermissions(c *gin.Context) {
	var path string
	err := c.Bind(&path)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	//userRoleID := ""
	//if c.Request.Header.Get("token") != "null" {
	//	accessDetails, err := h.helper.Token.GetAccessDetail(c)
	//	if h.helper.HTTP.HandleError(c, err) {
	//		return
	//	}
	//	userRoleID = accessDetails.UserRoleID
	//}
	//err = h.service.CheckPathPermissions(path, userRoleID)
	//if h.helper.HTTP.HandleError(c, err, http.StatusForbidden) {
	//	return
	//}
	c.JSON(http.StatusOK, nil)
}
