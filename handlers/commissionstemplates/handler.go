package commissionstemplates

import (
	"github.com/gin-gonic/gin"

	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.CommissionTemplate
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.Create(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	err := h.service.setQueryFilter(c)
	//if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
	//	return
	//}
	//userID, err := h.helper.Token.GetUserID(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusUnauthorized) {
		return
	}
	items, err := h.service.GetAll()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	err := h.service.setQueryFilter(c)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	id := c.Param("id")
	item, err := h.service.Get(id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(&id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.CommissionTemplate
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = h.service.Update(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetValueTypes(c *gin.Context) {
	items, err := h.service.GetValueTypes()
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetXlsx(c *gin.Context) {
	//var item models.CommissionTemplate
	//err := c.Bind(&item)
	//if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
	//	return
	//}
	//err = h.service.Update(&item)
	//if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
	//	return
	//}
	//excelDoc, err := h.helper.XLSX.CreateFile()
	//if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
	//	return
	//}
	//downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	//c.Header("Content-Description", "File Transfer")
	//c.Header("Content-Disposition", "attachment; filename="+downloadName)
	//c.Data(http.StatusOK, "application/octet-stream", excelDoc)
}
