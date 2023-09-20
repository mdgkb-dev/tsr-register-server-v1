package writers

import (
	"fmt"
	"time"
)

type DataWithStyle struct {
	Style Style
	Data  interface{}
}

type Style struct {
	Color string
}

func NewDataWithStyle(data interface{}, style Style) DataWithStyle {
	return DataWithStyle{Data: data, Style: style}
}

func GetNormalizedData(data interface{}) interface{} {
	switch d := data.(type) {
	case nil:
		return d
	case string:
		return d
	case int, uint:
		return d.(int)
	case float64:
		return fmt.Sprintf("%.2f", d)
	case float32:
		return fmt.Sprintf("%.2f", d)
	case *time.Time:
		return d.Format("02.01.2006")
	case time.Time:
		return d.Format("02.01.2006")
	case DataWithStyle:
	}

	return data
}
