package httpHelper

import (
	"encoding/json"
	"fmt"
	"time"
)

// SortModel model
type SortModel []map[string]string

// FilterModel model
type FilterModel struct {
	Table      *string   `json:"table"`
	Col        *string   `json:"col"`
	FilterType *string   `json:"filterType,omitempty"`
	Type       *DataType `json:"type,omitempty"`
	Operator   *Operator `json:"operator,omitempty"`
	Date1      time.Time `json:"date1,omitempty"`
	Date2      time.Time `json:"date2,omitempty"`

	Value1 string
	Value2 string
}

func (f *FilterModel) DatesToString() {
	f.Value1 = f.Date1.Format("2006-01-02")
	if f.IsBetween() {
		f.Value2 = f.Date2.Format("2006-01-02")
	}
	return
}

func (f *FilterModel) GetTableAndCol() string {
	return fmt.Sprintf("%s.%s", *f.Table, *f.Col)
}

func (f *FilterModel) IsUnary() bool {
	return *f.Operator == Eq || *f.Operator == Gt || *f.Operator == Ge
}

func (f *FilterModel) IsBetween() bool {
	return *f.Operator == Btw
}

// FilterModels model
type FilterModels []*FilterModel

type Operator string

const (
	Eq  Operator = "="
	Gt           = ">"
	Ge           = "<"
	Btw          = "between"
)

type DataType string

const (
	DateType   DataType = "date"
	NumberType          = "number"
	StringType          = "string"
)

// ParseJSONToFilterModel constructor
func ParseJSONToFilterModel(args string) (filterModel FilterModel, err error) {
	err = json.Unmarshal([]byte(args), &filterModel)
	if err != nil {
		return filterModel, err
	}
	return filterModel, err
}
