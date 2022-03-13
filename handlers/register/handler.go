package register

import (
	"github.com/gin-gonic/gin"
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Register
	err := c.Bind(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.Create(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	userId, err := models.GetUserID(c)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	items, err := h.service.GetAll(*userId)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	queryFilter, err := httpHelper.CreateQueryFilter(c)
	item, err := h.service.Get(queryFilter)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(&id)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Register
	err := c.Bind(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.Update(&item)
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetValueTypes(c *gin.Context) {
	items, err := h.service.GetValueTypes()
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetFlatXlsx(c *gin.Context) {
	//var item models.Register
	//err := c.Bind(&item)
	//if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
	//	return
	//}
	//err = h.service.Update(&item)
	//if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
	//	return
	//}
	//excelDoc, err := h.helper.XLSX.CreateFile()
	//if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
	//	return
	//}
	//downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	//c.Header("Content-Description", "File Transfer")
	//c.Header("Content-Disposition", "attachment; filename="+downloadName)
	//c.Data(http.StatusOK, "application/octet-stream", excelDoc)
}
