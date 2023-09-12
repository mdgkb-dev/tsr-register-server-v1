package models

import (
	"mdgkb/tsr-tegister-server-v1/helpers/xlsxhelper"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type XLSXWriter struct {
	bun.BaseModel `bun:"data_queries,alias:data_queries"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string        `json:"name"`
	Type          string        `bun:"type:register_query_type_enum" json:"type"`

	WithAge         bool `json:"withAge"`
	CountAverageAge bool `json:"countAverageAge"`
	Xl              *xlsxhelper.XlsxHelper
}

func (item *XLSXWriter) setStyle(xl *xlsxhelper.XlsxHelper) {
	xl.Cursor = 3
	//height := 6 + len(item.ResearchesPool.RegisterToPatient)
	//xl.SetBorder(height)
	//item.ResearchQueryGroups.writeAggregates(xl)
	//xl.AutofitAllColumns()
}

func (item *XLSXWriter) WriteFile(headers [][]interface{}, data [][]interface{}) ([]byte, error) {
	item.Xl = xlsxhelper.NewXlsxHelper()
	item.Xl.CreateFile()

	for lineN, line := range headers {
		for colN, colName := range line {
			item.Xl.WriteCell(lineN+1, colN, colName)
		}
	}

	headerLinesLen := len(headers)
	for lineN, line := range data {
		for colN, d := range line {
			item.Xl.WriteCell(headerLinesLen+lineN+1, colN, d)
		}
	}

	return item.Xl.WriteFile()
}
