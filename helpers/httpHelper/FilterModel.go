package httpHelper

import (
	"encoding/json"
	"fmt"
	"mdgkb/tsr-tegister-server-v1/helpers"
	"strings"
	"time"
)

// FilterModel model
type FilterModel struct {
	Table      *string `json:"table"`
	Col        *string `json:"col"`
	TableCol   string
	FilterType *string   `json:"filterType,omitempty"`
	Type       *DataType `json:"type,omitempty"`
	Operator   *Operator `json:"operator,omitempty"`
	Date1      time.Time `json:"date1,omitempty"`
	Date2      time.Time `json:"date2,omitempty"`

	Value1 string `json:"value1,omitempty"`
	Value2 string `json:"value2,omitempty"`

	Set []string `json:"set"`

	JoinTable   *string `json:"joinTable"`
	JoinTableFK *string `json:"joinTableFK"`
	JoinTablePK *string `json:"joinTablePK"`
}

func (f *FilterModel) DatesToString() {
	f.Value1 = fmt.Sprintf("'%s'", f.Date1.Format("2006-01-02"))
	if f.IsBetween() {
		f.Value2 = fmt.Sprintf("'%s'", f.Date2.Format("2006-01-02"))
	}
	f.TableCol = fmt.Sprintf("TO_CHAR(%s, 'YYYY-MM-DD')", f.TableCol)
	return
}

func (f *FilterModel) LikeToString() {
	//likeOperators := map[string]string{
	//	"contains":    "%%%s%%",
	//	"notContains": "%%%s%%",
	//	"startsWith":  "%s%%",
	//	"endsWith":    "%%%s",
	//}
	f.Value1 = fmt.Sprintf("'%%%s%%'", f.Value1)
	return
}

func (f *FilterModel) GetTableAndCol() string {
	return fmt.Sprintf("%s.%s", *f.Table, *f.Col)
}

func (f *FilterModel) SetTableAndCol() {
	f.TableCol = fmt.Sprintf("%s.%s", *f.Table, *f.Col)
}

func (f *FilterModel) Modify() {
	if *f.Type == StringType {
		if f.isNotBoolStringValues() {
			f.Translit()
			f.LikeToString()
			f.SetSmallCase()
		}
	}
	if *f.Type == DateType {
		f.DatesToString()
	}
}

func (f *FilterModel) isNotBoolStringValues() bool {
	return f.Value1 != "true" && f.Value1 != "false" && f.Value2 != "true" && f.Value2 != "false"
}

func (f *FilterModel) Translit() {
	f.Value1 = helpers.TranslitToRu(strings.ToLower(f.Value1))
	f.Value2 = helpers.TranslitToRu(strings.ToLower(f.Value2))
}

func (f *FilterModel) SetSmallCase() {
	f.TableCol = fmt.Sprintf("LOWER(%s)", f.TableCol)
	f.Value1 = fmt.Sprintf("LOWER(%s)", f.Value1)
	f.Value2 = fmt.Sprintf("LOWER(%s)", f.Value2)
}

func (f *FilterModel) GetJoinCondition() string {
	return fmt.Sprintf("%s.%s = %s.%s", *f.Table, *f.JoinTableFK, *f.JoinTable, *f.JoinTablePK)
}

func (f *FilterModel) IsUnary() bool {
	return *f.Operator == Eq || *f.Operator == Gt || *f.Operator == Ge || *f.Operator == Like
}

func (f *FilterModel) IsLike() bool {
	return *f.Operator == Like
}

func (f *FilterModel) IsBetween() bool {
	return *f.Operator == Btw
}

func (f *FilterModel) IsSet() bool {
	return *f.Operator == In
}

// FilterModels model
type FilterModels []*FilterModel

type Operator string

const (
	Eq   Operator = "="
	Gt            = ">"
	Ge            = "<"
	Btw           = "between"
	Like          = "like"
	In            = "in"
)

type DataType string

const (
	DateType   DataType = "date"
	NumberType          = "number"
	StringType          = "string"
	SetType             = "set"
)

// ParseJSONToFilterModel constructor
func ParseJSONToFilterModel(args string) (filterModel FilterModel, err error) {
	err = json.Unmarshal([]byte(args), &filterModel)
	if err != nil {
		return filterModel, err
	}
	return filterModel, err
}
