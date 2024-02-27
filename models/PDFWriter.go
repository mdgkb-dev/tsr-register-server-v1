package models

import (
	"mdgkb/tsr-tegister-server-v1/helpers/writers"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helpers/pdf"
	"github.com/uptrace/bun"
)

type PDFWriter struct {
	bun.BaseModel `bun:"data_queries,alias:data_queries"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string        `json:"name"`
	Type          string        `bun:"type:register_query_type_enum" json:"type"`

	WithAge         bool `json:"withAge"`
	CountAverageAge bool `json:"countAverageAge"`
	PDF             *pdf.PDF
}

func (item *PDFWriter) WriteFile(headers [][]interface{}, _ Agregator, data [][]interface{}) ([]byte, error) {
	for lineN := range headers {
		for colN := range headers[lineN] {
			headers[lineN][colN] = writers.GetNormalizedData(headers[lineN][colN])
		}
	}
	for lineN := range data {
		for colN := range data[lineN] {
			data[lineN][colN] = writers.GetNormalizedData(data[lineN][colN])
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
