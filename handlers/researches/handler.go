package researches

import (
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Research
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.Create(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	items, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) FTSP(c *gin.Context) {
	data, err := S.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, models.FTSPAnswer{Data: data, FTSP: *h.helper.SQL.ExtractFTSP(c.Request.Context())})
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	item, err := S.Get(c.Request.Context(), id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := S.Delete(c.Request.Context(), &id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Research
	_, err := h.helper.HTTP.GetForm(c, &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = S.Update(c.Request.Context(), &item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Xlsx(c *gin.Context) {
	//researchID := c.Param("research-id")
	//patientResearchID := c.Param("patient-id")
	//research, patient, err := S.GetResearchAndPatient(c, researchID, patientResearchID)
	//if h.helper.HTTP.HandleError(c, err) {
	//	return
	//}

	//researhQuery := models.DataQuery{}
	//researhQuery.Xl = xlsxhelper.NewXlsxHelper()
	//data, err := patient.GetXlsxData(research)
	//if h.helper.HTTP.HandleError(c, err) {
	//	return
	//}
	//file, err := researhQuery.WriteXlsxV2(research.GetHeaders(patient.Human.GetFullName()), data)

	//if h.helper.HTTP.HandleError(c, err) {
	//	return
	//}
	//downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	//c.Header("Content-Description", "File Transfer")
	//c.Header("Content-Disposition", "attachment; filename="+`"`+downloadName+`"`)
	//c.Data(http.StatusOK, "application/octet-stream", file)
}
