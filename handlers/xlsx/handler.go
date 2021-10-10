package xlsx

import (
	"github.com/gin-gonic/gin"
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
	"net/http"
	"time"
)

func (h *Handler) RegisterQuery(c *gin.Context) {
	//query, err := h.service.GetRegisterQuery(httpHelper.GetID(c))
	excelDoc, err := h.xlsxService.GetFile()
	if httpHelper.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", excelDoc)
	//c.Stream(http.StatusOK, "application/octet-stream", excelDoc)
}
