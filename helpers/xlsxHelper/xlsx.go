package xlsxHelper

import (
	"bytes"
	"fmt"
	"github.com/xuri/excelize/v2"
)

type IXlsxHelper interface {
	CreateFile() ([]byte, error)
	//GetFile() (*bytes.Reader, error)
}

type XlsxHelper struct{}

func NewXlsxHelper() *XlsxHelper {
	return &XlsxHelper{}
}

// Write func
func (x *XlsxHelper) writeData(file *excelize.File, data []map[string]interface{}) error {
	if len(data) == 0 {
		return nil
	}
	keys := make([]int, 0, len(data[0]))
	for k := range data {
		keys = append(keys, k)
	}

	err := file.SetSheetRow("Sheet1", "A1", &keys)
	if err != nil {
		return err
	}

	for i, mapa := range data {
		if i == 0 {
			continue
		}
		values := make([]interface{}, 0)
		for _, value := range mapa {
			values = append(values, value)
			err := file.SetCellValue("Sheet1", "A1", value)
			if err != nil {
				return err
			}
		}
		row := fmt.Sprintf("A%d", i)
		err := file.SetSheetRow("Sheet1", row, &values)
		if err != nil {
			return err
		}
	}
	return nil
}

func (x *XlsxHelper) CreateFile(data []map[string]interface{}) ([]byte, error) {
	f := excelize.NewFile()
	err := x.writeData(f, data)
	if err != nil {
		return nil, err
	}
	return x.write(f)
}

// Write func
func (x *XlsxHelper) write(file *excelize.File) ([]byte, error) {
	var b bytes.Buffer
	err := file.Write(&b)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
