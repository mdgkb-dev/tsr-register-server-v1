package xlsxHelper

import (
	"bytes"

	"github.com/xuri/excelize/v2"
)

type IXlsxHelper interface {
	CreateFile() ([]byte, error)
	//GetFile() (*bytes.Reader, error)
}

type XlsxHelper struct{}

func CreateXlsxHelper() *XlsxHelper {
	return &XlsxHelper{}
}

func (x *XlsxHelper) CreateFile() ([]byte, error) {
	f := excelize.NewFile()

	_ = f.SetCellValue("Sheet2", "A2", "Hello world.")
	_ = f.SetCellValue("Sheet1", "B2", 100)

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
