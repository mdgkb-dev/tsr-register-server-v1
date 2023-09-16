package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/pdfHelper"
	"github.com/uptrace/bun"
)

type PDFWriter struct {
	bun.BaseModel `bun:"data_queries,alias:data_queries"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string        `json:"name"`
	Type          string        `bun:"type:register_query_type_enum" json:"type"`

	WithAge         bool `json:"withAge"`
	CountAverageAge bool `json:"countAverageAge"`
	PDF             *pdfHelper.PDFHelper
}

func NormalizeExportData(data interface{}) interface{} {
	switch d := data.(type) {
	case string:
		return d
	case int, uint:
		return d.(int)
	case float64:
		return fmt.Sprintf("%.2f", d)
	case float32:
		fmt.Print(d)
		return fmt.Sprintf("%.2f", d)
	case *time.Time:
		return d.Format("02.01.2006")
	case time.Time:
		return d.Format("02.01.2006")
	}
	return data
}

func (item *PDFWriter) WriteFile(headers [][]interface{}, data [][]interface{}) ([]byte, error) {

	for lineN := range headers {
		for colN := range headers[lineN] {
			headers[lineN][colN] = NormalizeExportData(headers[lineN][colN])
		}
	}
	for lineN := range data {
		for colN := range data[lineN] {
			data[lineN][colN] = NormalizeExportData(data[lineN][colN])
		}
	}
	return item.PDF.GeneratePDF("patientResearch", struct {
		Headers [][]interface{}
		D       [][]interface{}
	}{
		Headers: headers,
		D:       data,
	})
}
