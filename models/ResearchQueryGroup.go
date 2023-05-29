package models

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/helpers/xlsxhelper"
	"sort"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResearchQueryGroup struct {
	bun.BaseModel   `bun:"research_query_groups,alias:research_query_groups"`
	ID              uuid.NullUUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name            string         `json:"name"`
	ResearchQueryID uuid.NullUUID  `bun:"type:uuid" json:"researchQueryId"`
	ResearchQuery   *ResearchQuery `bun:"rel:belongs-to" json:"researchQuery"`

	ResearchID uuid.UUID `bun:"type:uuid" json:"researchId"`
	Research   *Research `bun:"rel:belongs-to" json:"research"`

	ResearchQueryGroupQuestions ResearchQueryGroupQuestions `bun:"rel:has-many" json:"registerQueryGroupProperties"`

	Order int `bun:"item_order" json:"order"`

	AggregateType        AggregateType `json:"aggregateType"`
	AggregatedProperties Questions     `bun:"rel:has-many" json:"aggregatedProperties"`

	CountSum        bool `json:"countSum"`
	CountPercentage bool `json:"countPercentage"`

	Sum                      uint                     `bun:"-" json:"sum"`
	Percentage               uint                     `bun:"-" json:"percentage"`
	AggregatedValues         map[string]float64       `bun:"-" json:"aggregatedValues"`
	RegisterQueryPercentages ResearchQueryPercentages `bun:"-" `

	PatientIndex int `bun:"-"`
}

type ResearchQueryPercentage struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

type ResearchQueryPercentages []*ResearchQueryPercentage

type ResearchQueryGroups []*ResearchQueryGroup

func (item *ResearchQueryGroup) GetResultFromData(prop *Question, propertyIndex int) string {
	return item.getAggregatedData(prop, propertyIndex)
}

func (item *ResearchQueryGroup) getAggregatedData(question *Question, propertyIndex int) string {
	if item.AggregateType == AggregateNone {
		if question.ValueType.IsSet() {
			return question.AnswerVariants.Include(item.Research.ResearchResults[item.PatientIndex].Answers)
		}
		if len(item.Research.ResearchResults[item.PatientIndex].Answers) > 0 {
			if propertyIndex < len(item.Research.ResearchResults[item.PatientIndex].Answers) {
				return item.Research.ResearchResults[item.PatientIndex].Answers[propertyIndex].GetData(question)
			}
			return No
		}
	}
	if item.AggregateType == AggregateExisting {
		if item.Research.ResearchResults != nil && item.PatientIndex < len(item.Research.ResearchResults) {
			return item.Research.ResearchResults[item.PatientIndex].GetAggregateExistingData()
		}
	}
	return ""
}

func (item *ResearchQueryGroup) GetAggregatedPercentage() {
	sum := float64(0)
	for k, v := range item.AggregatedValues {
		sum += v
		item.RegisterQueryPercentages = append(item.RegisterQueryPercentages, &ResearchQueryPercentage{k, v})
	}
	sort.Slice(item.RegisterQueryPercentages, func(i, j int) bool {
		return item.RegisterQueryPercentages[i].Value > item.RegisterQueryPercentages[j].Value
	})
}

func (items ResearchQueryGroups) writeXlsxHeader(xl *xlsxhelper.XlsxHelper) {
	for i := range items {
		items[i].writeXlsxHeader(xl)
	}
}

func (item *ResearchQueryGroup) writeXlsxHeader(xl *xlsxhelper.XlsxHelper) {
	xl.WriteString(1, xl.Cursor, &[]string{item.Name})
	if item.AggregateType == AggregateNone {
		item.ResearchQueryGroupQuestions.writeXlsxHeader(xl)
	} else {
		xl.WriteString(1, xl.Cursor, &[]string{item.Name})
		xl.Cursor++
	}
}

func (items ResearchQueryGroups) writeXlsxData(xl *xlsxhelper.XlsxHelper, id uuid.NullUUID) {
	for i := range items {
		writeEmpty := false
		if items[i].PatientIndex >= len(items[i].Research.ResearchResults) {
			return
		}
		if items[i].Research.ResearchResults[items[i].PatientIndex] == nil {
			return
		}
		//if items[i].Research.ResearchResults[items[i].PatientIndex].PatientResearch.PatientID != id {
		//	writeEmpty = true
		//}
		items[i].writeXlsxData(xl, writeEmpty)
		if !writeEmpty {
			items[i].PatientIndex++
		}
	}
}

func (item *ResearchQueryGroup) writeXlsxData(xl *xlsxhelper.XlsxHelper, writeEmpty bool) {
	if item.AggregateType == AggregateNone {
		item.ResearchQueryGroupQuestions.writeXlsxData(xl, item, writeEmpty)
	}
	if item.AggregateType == AggregateExisting {
		res := item.GetResultFromData(nil, 0)
		if writeEmpty {
			res = NoData
		}
		str := fmt.Sprintf("%v", res)
		item.writePercentage(str)
		xl.Data = append(xl.Data, str)
	}
}

func (item *ResearchQueryGroup) writePercentage(key string) {
	_, ok := item.AggregatedValues[key]
	if ok {
		item.AggregatedValues[key]++
	} else {
		item.AggregatedValues[key] = 1
	}
}

func (items ResearchQueryGroups) writeAggregates(xl *xlsxhelper.XlsxHelper) {
	for i := range items {
		items[i].writeAggregates(xl)
	}
}

func (item *ResearchQueryGroup) writeAggregates(xl *xlsxhelper.XlsxHelper) {
	if item.AggregateType == AggregateNone {
		item.ResearchQueryGroupQuestions.writeAggregates(xl)
	}
	if item.AggregateType == AggregateExisting {
		item.AggregateType.WriteAggregatedValues(xl, item.AggregatedValues)
	}
}
