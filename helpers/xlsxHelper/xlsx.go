package xlsxHelper

import (
	"bytes"
	"fmt"
	"github.com/xuri/excelize/v2"
)

type IXlsxHelper interface {
	CreateFile() ([]byte, error)
}

type XlsxHelper struct {
	file        *excelize.File
	keys        []string
	data        []map[string]interface{}
	HeaderCells []string

	style *DefaultStyle
}

func (x *XlsxHelper) CreateFile(keys []string, data []map[string]interface{}) ([]byte, error) {
	x.file = excelize.NewFile()
	x.initXlsxData(keys, data)
	err := x.writeHeader()
	if err != nil {
		return nil, err
	}
	err = x.writeData()
	if err != nil {
		return nil, err
	}
	return x.writeFile()
}

func (x *XlsxHelper) initXlsxData(keys []string, data []map[string]interface{}) {
	x.keys = keys
	x.data = data
}

func NewXlsxHelper() *XlsxHelper {
	return &XlsxHelper{}
}

func (x *XlsxHelper) stylingHeader() error {
	style, err := x.file.NewStyle(&excelize.Style{
		Fill: excelize.Fill{Type: "pattern", Color: []string{"#D0D0D0"}, Pattern: 1},
	})
	if err != nil {
		return err
	}
	err = x.file.SetColWidth("Sheet1", "A", "BA", 30)
	if err != nil {
		return err
	}
	err = x.file.SetCellStyle("Sheet1", "A1", "BA1", style)
	if err != nil {
		return err
	}
	return err
}

func (x *XlsxHelper) writeHeader() error {
	err := x.file.SetSheetRow("Sheet1", "A1", &x.keys)
	if err != nil {
		return err
	}
	err = x.stylingHeader()
	if err != nil {
		return err
	}
	return err
}

func (x *XlsxHelper) writeData() error {
	if len(x.data) == 0 {
		return nil
	}
	for i, mapa := range x.data {
		values := make([]interface{}, 0)
		for _, k := range x.keys {
			values = append(values, mapa[k])
		}
		row := fmt.Sprintf("A%d", i+2)
		//fmt.Println(row)
		err := x.file.SetSheetRow("Sheet1", row, &values)
		if err != nil {
			return err
		}
	}
	return nil
}

// Write func
func (x *XlsxHelper) writeFile() ([]byte, error) {
	var b bytes.Buffer
	err := x.file.Write(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
