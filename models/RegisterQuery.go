package models

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/helpers/xlsxhelper"
	"strconv"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterQuery struct {
	bun.BaseModel                            `bun:"register_queries,alias:register_queries"`
	ID                                       uuid.UUID                          `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name                                     string                             `json:"name"`
	Type                                     string                             `bun:"type:register_query_type_enum" json:"type"`
	Register                                 *Register                          `bun:"rel:belongs-to" json:"register"`
	RegisterID                               uuid.UUID                          `bun:"type:uuid" json:"registerId"`
	RegisterQueryToRegisterProperty          []*RegisterQueryToRegisterProperty `bun:"rel:has-many" json:"registerQueryToRegisterProperty"`
	RegisterQueryToRegisterPropertyForDelete []string                           `bun:"-" json:"registerQueryToRegisterPropertyForDelete"`
	Keys                                     []string                           `bun:"-" json:"keys"`
	Key                                      string                             `json:"key"`

	RegisterQueryGroups RegisterQueryGroups `bun:"rel:has-many" json:"registerQueryGroups"`

	WithAge         bool `json:"withAge"`
	CountAverageAge bool `json:"countAverageAge"`
}

type RegisterQueries []*RegisterQuery

func (item *RegisterQuery) SetIDForChildren() {
	if len(item.RegisterQueryToRegisterProperty) == 0 {
		return
	}

	for i := range item.RegisterQueryToRegisterProperty {
		item.RegisterQueryToRegisterProperty[i].RegisterQueryID = item.ID
	}
}

func (item *RegisterQuery) writeHeaderStandardCols(xl *xlsxhelper.XlsxHelper) {
	standardCols := []string{"№", "ФИО"}
	if item.WithAge {
		standardCols = append(standardCols, "Возраст")
	}
	for i := range standardCols {
		colName := xl.GetCol(i)
		xl.MergeArea(fmt.Sprintf("%s1", colName), fmt.Sprintf("%s3", colName))
	}
	xl.WriteString(1, 0, &standardCols)
	xl.Cursor = len(standardCols)
}

func (item *RegisterQuery) writeXlsxHeader(xl *xlsxhelper.XlsxHelper) {
	item.writeHeaderStandardCols(xl)
	item.RegisterQueryGroups.writeXlsxHeader(xl)
}

func (item *RegisterQuery) WriteXlsx(xl *xlsxhelper.XlsxHelper) ([]byte, error) {
	xl.CreateFile()
	item.writeXlsxHeader(xl)
	if xl.Err != nil {
		return nil, xl.Err
	}
	item.writeData(xl)
	if xl.Err != nil {
		return nil, xl.Err
	}
	item.writeAggregates(xl)
	if xl.Err != nil {
		return nil, xl.Err
	}
	item.setStyle(xl)
	if xl.Err != nil {
		return nil, xl.Err
	}
	return xl.WriteFile()
}

func (item *RegisterQuery) writeData(xl *xlsxhelper.XlsxHelper) {
	for groupNum := range item.RegisterQueryGroups {
		item.RegisterQueryGroups[groupNum].AggregatedValues = make(map[string]float64)
		for propNum := range item.RegisterQueryGroups[groupNum].RegisterQueryGroupProperties {
			item.RegisterQueryGroups[groupNum].RegisterQueryGroupProperties[propNum].AggregatedValues = make(map[string]float64)
			//for setNum := range item.RegisterQueryGroups[groupNum].RegisterQueryGroupProperties[propNum].ResearchResult.RegisterPropertySets {
			//	item.RegisterQueryGroups[groupNum].RegisterQueryGroupProperties[propNum].ResearchResult.RegisterPropertySets[setNum].AggregatedValues = make(map[string]float64)
			//}
			for radioNum := range item.RegisterQueryGroups[groupNum].RegisterQueryGroupProperties[propNum].RegisterProperty.AnswerVariants {
				item.RegisterQueryGroups[groupNum].RegisterQueryGroupProperties[propNum].RegisterProperty.AnswerVariants[radioNum].AggregatedValues = make(map[string]float64)
			}
		}
	}

	for patientNum, _ := range item.Register.RegisterToPatient {
		//xl.Data = append(xl.Data, strconv.Itoa(patientNum+1), registerToPatient.Patient.Human.GetFullName())
		//if item.WithAge {
		//	xl.Data = append(xl.Data, strconv.Itoa(registerToPatient.Patient.Human.GetAge()))
		//}
		//item.RegisterQueryGroups.writeXlsxData(xl, registerToPatient.PatientID)
		xl.WriteString(4+patientNum, 0, &xl.Data)
		xl.Data = []string{}
	}
}

func (item *RegisterQuery) writeAggregates(xl *xlsxhelper.XlsxHelper) {
	xl.StrCursor = 4 + len(item.Register.RegisterToPatient)
	if item.WithAge {
		xl.WriteString(xl.StrCursor, 2, &[]string{strconv.Itoa(item.Register.GetPatientsAverageAge())})
	}
	xl.Cursor = 3
	item.RegisterQueryGroups.writeAggregates(xl)
}

func (item *RegisterQuery) setStyle(xl *xlsxhelper.XlsxHelper) {
	xl.Cursor = 3
	height := 6 + len(item.Register.RegisterToPatient)
	xl.SetBorder(height)
	item.RegisterQueryGroups.writeAggregates(xl)
	//xl.AutofitAllColumns()
}
