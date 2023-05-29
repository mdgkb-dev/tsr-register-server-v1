package researchquery

import (
	"mdgkb/tsr-tegister-server-v1/helpers/xlsxhelper"
	"time"

	"github.com/gin-gonic/gin"

	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"
)

func (h *Handler) Create(c *gin.Context) {
	var query models.ResearchQuery
	err := c.Bind(&query)

	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	err = h.service.Create(&query)

	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, query)
}

func (h *Handler) GetAll(c *gin.Context) {
	queries, err := h.service.GetAll()

	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, queries)
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	query, err := h.service.Get(id)

	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, query)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var query models.ResearchQuery
	err := c.Bind(&query)

	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	err = h.service.Update(&query)

	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}

	c.JSON(http.StatusOK, query)
}

func (h *Handler) Execute(c *gin.Context) {
	id := c.Param("id")
	registerQuery, err := h.service.Get(id)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	//err = h.service.Execute(registerQuery)
	//if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
	//	return
	//}
	file, err := registerQuery.WriteXlsx(xlsxhelper.NewXlsxHelper())
	//x :=
	//x.WriteHeader()
	//file, err := xlsxhelper.NewXlsxHelper().CreateFile(registerQuery.Keys, registerQuery.Data)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+`"`+downloadName+`"`)
	c.Data(http.StatusOK, "application/octet-stream", file)
	//c.JSON(http.StatusOK, registerQuery)
}
