package models

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/helpers/xlsxhelper"
	"strconv"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResearchQuery struct {
	bun.BaseModel            `bun:"research_queries,alias:research_queries"`
	ID                       uuid.NullUUID            `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name                     string                   `json:"name"`
	Type                     string                   `bun:"type:register_query_type_enum" json:"type"`
	ResearchesPool           *ResearchesPool          `bun:"rel:belongs-to" json:"researchesPool"`
	ResearchesPoolID         uuid.NullUUID            `bun:"type:uuid" json:"researchesPoolId"`
	ResearchQueriesQuestions ResearchQueriesQuestions `bun:"rel:has-many" json:"researchQueriesQuestions"`
	Keys                     []string                 `bun:"-" json:"keys"`
	Key                      string                   `json:"key"`

	ResearchQueryGroups ResearchQueryGroups `bun:"rel:has-many" json:"registerQueryGroups"`

	WithAge         bool `json:"withAge"`
	CountAverageAge bool `json:"countAverageAge"`
	Xl              *xlsxhelper.XlsxHelper
}

type ResearchQueries []*ResearchQuery
type ResearchQueriesWithCount struct {
	ResearchQueries ResearchQueries `json:"items"`
	Count           int             `json:"count"`
}

func (item *ResearchQuery) SetIDForChildren() {
	if len(item.ResearchQueriesQuestions) == 0 {
		return
	}

	for i := range item.ResearchQueriesQuestions {
		item.ResearchQueriesQuestions[i].ResearchQueryID = item.ID
	}
}

func (item *ResearchQuery) writeHeaderStandardCols(xl *xlsxhelper.XlsxHelper) {
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

func (item *ResearchQuery) writeXlsxHeader(xl *xlsxhelper.XlsxHelper) {
	item.writeHeaderStandardCols(xl)
	item.ResearchQueryGroups.writeXlsxHeader(xl)
}

func (item *ResearchQuery) WriteXlsx(xl *xlsxhelper.XlsxHelper) ([]byte, error) {
	xl.CreateFile()
	item.writeXlsxHeader(xl)
	if xl.Err != nil {
		return nil, xl.Err
	}

	item.writeData(xl)
	if xl.Err != nil {
		return nil, xl.Err
	}
	//item.writeAggregates(xl)
	//if xl.Err != nil {
	//	return nil, xl.Err
	//}
	//item.setStyle(xl)
	//if xl.Err != nil {
	//	return nil, xl.Err
	//}
	return xl.WriteFile()
}

func (item *ResearchQuery) writeData(xl *xlsxhelper.XlsxHelper) {
	for groupNum := range item.ResearchQueryGroups {
		item.ResearchQueryGroups[groupNum].AggregatedValues = make(map[string]float64)
		for propNum := range item.ResearchQueryGroups[groupNum].ResearchQueryGroupQuestions {
			item.ResearchQueryGroups[groupNum].ResearchQueryGroupQuestions[propNum].AggregatedValues = make(map[string]float64)
			//for radioNum := range item.ResearchQueryGroups[groupNum].ResearchQueryGroupQuestions[propNum].Question.AnswerVariants {
			//	item.ResearchQueryGroups[groupNum].ResearchQueryGroupQuestions[propNum].Question.AnswerVariants[radioNum].AggregatedValues = make(map[string]float64)
			//}
		}
	}
	for patientNum, patientResearchPool := range item.ResearchesPool.PatientsResearchesPools {
		if patientResearchPool.Patient == nil {
			continue
		}
		xl.Data = append(xl.Data, strconv.Itoa(patientNum+1), patientResearchPool.Patient.Human.GetFullName())
		if item.WithAge {
			xl.Data = append(xl.Data, patientResearchPool.Patient.Human.GetFormattedDateBirth())
		}
		item.ResearchQueryGroups.writeXlsxData(xl, patientResearchPool.PatientID)
		xl.WriteString(4+patientNum, 0, &xl.Data)
		//fmt.Println(len(xl.Data))
		xl.Data = []string{}
	}
}

func (item *ResearchQuery) writeAggregates(xl *xlsxhelper.XlsxHelper) {
	xl.StrCursor = 4 + len(item.ResearchesPool.PatientsResearchesPools)
	if item.WithAge {
		xl.WriteString(xl.StrCursor, 2, &[]string{strconv.Itoa(item.ResearchesPool.GetPatientsAverageAge())})
	}
	xl.Cursor = 3
	item.ResearchQueryGroups.writeAggregates(xl)
}

func (item *ResearchQuery) setStyle(xl *xlsxhelper.XlsxHelper) {
	xl.Cursor = 3
	//height := 6 + len(item.ResearchesPool.RegisterToPatient)
	//xl.SetBorder(height)
	//item.ResearchQueryGroups.writeAggregates(xl)
	//xl.AutofitAllColumns()
}

func (item *ResearchQuery) WriteXlsxV2(headers [][]string, data [][]string) ([]byte, error) {
	item.Xl.CreateFile()

	for lineN, line := range headers {
		for colN, colName := range line {
			fmt.Println(colName)
			item.Xl.WriteCell(lineN+1, colN, colName)
		}
	}

	headerLinesLen := len(headers)
	for lineN, line := range data {
		for colN, d := range line {
			item.Xl.WriteCell(headerLinesLen+lineN+1, colN, d)
		}
	}

	return item.Xl.WriteFile()
}
