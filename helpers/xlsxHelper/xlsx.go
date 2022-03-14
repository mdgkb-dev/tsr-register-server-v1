package xlsxHelper

import (
	"bytes"
	"fmt"
	"github.com/xuri/excelize/v2"
	"sort"
)

type IXlsxHelper interface {
	CreateFile() ([]byte, error)
	//GetFile() (*bytes.Reader, error)
}

type XlsxHelper struct {
	HeaderCells []string
}

func NewXlsxHelper() *XlsxHelper {
	return &XlsxHelper{}
}

// Write func
func (x *XlsxHelper) writeData(file *excelize.File, data []map[string]interface{}) error {
	if len(data) == 0 {
		return nil
	}
	for i, mapa := range data {
		keys := make([]string, 0)
		for k, _ := range mapa {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Println(k, mapa[k])
		}

		values := make([]interface{}, 0)
		for _, k := range keys {
			if i == 0 {
				err := file.SetSheetRow("Sheet1", "A1", &keys)
				if err != nil {
					return err
				}
				continue
			}
			values = append(values, mapa[k])
			//err := file.SetCellValue("Sheet1", "A1", value)
			//if err != nil {
			//	return err
			//}
		}
		row := fmt.Sprintf("A%d", i+1)
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
