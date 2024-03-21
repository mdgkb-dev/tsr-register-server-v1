package dataexport

import (
	"fmt"
	"net/http"
	"time"

	"mdgkb/tsr-tegister-server-v1/handlers/patients"
	"mdgkb/tsr-tegister-server-v1/handlers/researches"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/gin-gonic/gin"
)

type ExportOpts struct {
	ExportOptions models.ExportOptions `json:"exportOptions"`
}

func (h *Handler) Export(c *gin.Context) {
	opts := ExportOpts{}
	// exportOptions :=
	_, err := h.helper.HTTP.GetForm(c, &opts)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	fmt.Println(opts)
	exportOptions := opts.ExportOptions

	fmt.Println(2)
	researchesExport := models.ResearchesExport{}
	err = exportOptions.Parse(&researchesExport)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	fmt.Println(3)
	researchesForExport, err := researches.R.GetForExport(c.Request.Context(), researchesExport.IDPool)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	fmt.Println(4)
	headers := researchesForExport.GetExportData()
	patientsExport := models.PatientsExport{}
	err = exportOptions.Parse(&patientsExport)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	fmt.Println(5)
	patientsForExport, err := patients.R.GetForExport(c.Request.Context(), patientsExport.IDPool)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	exportData, agregator, err := patientsForExport.Patients.GetExportData(researchesForExport)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	exporter := exportOptions.ExportType.GetExporter(h.helper)
	file, err := exporter.WriteFile(headers, agregator, exportData)
	if h.helper.HTTP.HandleError(c, err) {
		return
	}
	ext := ".pdf"
	if exportOptions.ExportType == models.ExportTypeXLSX {
		ext = ".xlsx"
	}
	fmt.Println(7)
	downloadName := time.Now().UTC().Format("data-20060102150405" + ext)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+`"`+downloadName+`"`)
	c.Data(http.StatusOK, "application/octet-stream", file)
}
