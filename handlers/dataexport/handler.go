package dataexport

import (
	"encoding/json"
	"fmt"
	"mdgkb/tsr-tegister-server-v1/handlers/patients"
	"mdgkb/tsr-tegister-server-v1/handlers/researches"
	"mdgkb/tsr-tegister-server-v1/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Export(c *gin.Context) {
	exportOptions := models.ExportOptions{}
	err := json.Unmarshal([]byte(c.Query("exportOptions")), &exportOptions)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	researchesExport := models.ResearchesExport{}
	err = exportOptions.Parse(&researchesExport)
	fmt.Println(researchesExport)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	researchesForExport, err := researches.R.GetForExport(c.Request.Context(), researchesExport.IDPool)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	headers := researchesForExport.GetExportData()
	patientsExport := models.PatientsExport{}
	err = exportOptions.Parse(&patientsExport)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	patientsForExport, err := patients.R.GetForExport(c.Request.Context(), patientsExport.IDPool)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	fmt.Println(patientsExport)
	exportData, err := patientsForExport.Patients.GetExportData(researchesForExport)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	exporter := exportOptions.ExportType.GetExporter()
	file, err := exporter.WriteFile(headers, exportData)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}

	downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+`"`+downloadName+`"`)
	c.Data(http.StatusOK, "application/octet-stream", file)
}
