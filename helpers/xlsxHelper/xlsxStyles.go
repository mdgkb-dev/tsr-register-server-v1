package xlsxHelper

import (
	"github.com/xuri/excelize/v2"
)

type DefaultStyle struct {
	header excelize.Style
	data   excelize.Style
}

func getDefaultStyle() *DefaultStyle {
	//defaultStyle := DefaultStyle{}
	//defaultStyle.initDefaultHeaderStyle()
	return getDefaultStyle()
}

func (i *DefaultStyle) initDefaultHeaderStyle() {
	//gray := "#D0D0D0"
	//i.header =
}
