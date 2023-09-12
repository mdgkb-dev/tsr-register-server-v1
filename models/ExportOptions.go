package models

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
	WriteFile([][]interface{}, [][]interface{}) ([]byte, error)
}

func (item ExportType) GetExporter() FileWriter {
	switch item {
	case ExportTypeXLSX:
		return &XLSXWriter{}
	case ExportTypeDOCX:
		return nil
	case ExportTypePDF:
		return nil
	default:
		return nil
	}
}

// Может выгружать:
// 1) Данные по конкретному исследованию у одного пациента
// 2) Данные по всему пациенту по всем исследованиям
// 3) Данные по всем пациентам и без исследований, и по всем исследованиям
// 4) Не знает о том, что нужно выгружать, это задаётся опциями запроса, умеет только писать

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
