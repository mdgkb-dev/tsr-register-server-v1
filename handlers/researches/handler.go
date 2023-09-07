package researches

import (
	"mdgkb/tsr-tegister-server-v1/helpers/xlsxhelper"
	"time"

	"github.com/gin-gonic/gin"

	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"
)

func (h *Handler) Create(c *gin.Context) {
	var item models.Research
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.Create(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetAll(c *gin.Context) {
	items, err := h.service.GetAll(c.Request.Context())
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Get(c *gin.Context) {
	err := h.service.setQueryFilter(c)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	id := c.Param("id")
	item, err := h.service.Get(id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(&id)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) Update(c *gin.Context) {
	var item models.Research
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	err = h.service.Update(&item)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) GetValueTypes(c *gin.Context) {
	items, err := h.service.GetValueTypes()
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *Handler) Xlsx(c *gin.Context) {
	researchID := c.Param("research-id")
	patientResearchID := c.Param("patient-id")
	research, patient, err := h.service.GetResearchAndPatient(c, researchID, patientResearchID)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	researhQuery := models.ResearchQuery{}
	researhQuery.Xl = xlsxhelper.NewXlsxHelper()
	data, err := patient.GetXlsxData(research)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	file, err := researhQuery.WriteXlsxV2(research.GetHeaders(patient.Human.GetFullName()), data)

	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+`"`+downloadName+`"`)
	c.Data(http.StatusOK, "application/octet-stream", file)
}
