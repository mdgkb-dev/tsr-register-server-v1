package models

import (
	"fmt"

	"github.com/pro-assistance/pro-assister/helper"
)

type ExportOptions struct {
	ExportType ExportType `json:"exportType"`

	Options map[string]map[string]interface{} `json:"options"`
}

type ExportType string

const (
	ExportTypeXLSX ExportType = "xlsx"
	ExportTypeDOCX ExportType = "docx"
	ExportTypePDF  ExportType = "pdf"
)

type FileWriter interface {
	WriteFile([][]interface{}, Agregator, [][]interface{}) ([]byte, error)
}

type Agregator struct {
	Count []int
	Sums  []float64
}

func (item *Agregator) GetAverage(i int) interface{} {
	if item.Count[i] == 0 {
		return ""
	}
	fmt.Println("res", item.Sums[i], item.Count[i])
	return item.Sums[i] / float64(item.Count[i])
}

func NewAgregator(length int) Agregator {
	return Agregator{Count: make([]int, length), Sums: make([]float64, length)}
}

func (item ExportType) GetExporter(helper *helper.Helper) FileWriter {
	switch item {
	case ExportTypeXLSX:
		return &XLSXWriter{}
	case ExportTypeDOCX:
		return nil
	case ExportTypePDF:
		writer := &PDFWriter{}
		writer.PDF = helper.PDF
		return writer
	default:
		return nil
	}
}

type OptionsParser interface {
	ParseExportOptions(map[string]map[string]interface{}) error
}

func (item *ExportOptions) Parse(parsers ...OptionsParser) error {
	var err error
	for i := range parsers {
		err = parsers[i].ParseExportOptions(item.Options)
		if err != nil {
			break
		}
	}
	return err
}
