package models

import (
	"mdgkb/tsr-tegister-server-v1/helpers/xlsxhelper"
	"sort"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResearchQueryGroupQuestion struct {
	bun.BaseModel        `bun:"register_query_group_properties,alias:register_query_group_properties"`
	ID                   uuid.UUID           `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	ResearchQueryGroupID uuid.NullUUID       `bun:"type:uuid" json:"researchQueryGroupId"`
	ResearchQueryGroup   *ResearchQueryGroup `bun:"rel:belongs-to" json:"researchQueryGroup"`

	ResearchResultID uuid.UUID       `bun:"type:uuid" json:"researchResultId"`
	ResearchResult   *ResearchResult `bun:"rel:belongs-to" json:"researchResult"`

	Question   *Question `bun:"rel:belongs-to" json:"question"`
	QuestionID uuid.UUID `bun:"type:uuid" json:"questionId"`

	Order                    int                      `bun:"item_order" json:"order"`
	AggregatedValues         map[string]float64       `bun:"-" json:"aggregatedValues"`
	ResearchQueryPercentages ResearchQueryPercentages `bun:"-" `
	EveryRadioSet            bool                     `json:"everyRadioSet"`

	AggregateType AggregateType `json:"aggregateType"`
}

type ResearchQueryGroupQuestions []*ResearchQueryGroupQuestion

func (item *ResearchQueryGroupQuestion) GetAggregatedPercentage() {
	sum := float64(0)
	for k, v := range item.AggregatedValues {
		sum += v
		item.ResearchQueryPercentages = append(item.ResearchQueryPercentages, &ResearchQueryPercentage{k, v})
	}
	sort.Slice(item.ResearchQueryPercentages, func(i, j int) bool {
		return item.ResearchQueryPercentages[i].Value > item.ResearchQueryPercentages[j].Value
	})
}

func (items ResearchQueryGroupQuestions) writeXlsxHeader(xl *xlsxhelper.XlsxHelper) {
	for i := range items {
		items[i].writeXlsxHeader(xl)
	}
}

func (item *ResearchQueryGroupQuestion) writeXlsxHeader(xl *xlsxhelper.XlsxHelper) {
	//if item.EveryRadioSet {
	//	xl.WriteString(2, xl.Cursor, &[]string{item.ResearchResult.Name})
	//	if item.ResearchResult.ValueType.IsSet() {
	//		for _, setItem := range item.ResearchResult.RegisterPropertySets {
	//			xl.WriteString(3, xl.Cursor, &[]string{setItem.Name})
	//			xl.Cursor++
	//			for _, other := range setItem.RegisterPropertyOthers {
	//				xl.WriteString(3, xl.Cursor, &[]string{other.Name})
	//				xl.Cursor++
	//			}
	//		}
	//	}
	//	if item.ResearchResult.ValueType.IsRadio() {
	//		for _, radioItem := range item.ResearchResult.AnswerVariants {
	//			xl.WriteString(3, xl.Cursor, &[]string{radioItem.Name})
	//			xl.Cursor++
	//			for _, other := range radioItem.RegisterPropertyOthers {
	//				xl.WriteString(3, xl.Cursor, &[]string{other.Name})
	//				xl.Cursor++
	//			}
	//		}
	//	}
	//} else {
	//	xl.WriteString(2, xl.Cursor, &[]string{item.ResearchResult.Name})
	//	xl.Cursor++
	//}
}

func (items ResearchQueryGroupQuestions) writeXlsxData(xl *xlsxhelper.XlsxHelper, g *ResearchQueryGroup, writeEmpty bool) {
	for i := range items {
		items[i].writeXlsxData(xl, g, i, writeEmpty)
	}
}

func (item *ResearchQueryGroupQuestion) writeXlsxData(xl *xlsxhelper.XlsxHelper, g *ResearchQueryGroup, propNum int, writeEmpty bool) {
	//if item.EveryRadioSet {
	//	if item.ResearchResult.ValueType.IsSet() {
	//		for i := range item.ResearchResult.RegisterPropertySets {
	//			v := No
	//			if g.Research.RegisterGroupsToPatients[g.PatientIndex].RegisterPropertySetToPatient.Include(item.ResearchResult.RegisterPropertySets[i].ID) {
	//				v = Yes
	//			}
	//			if writeEmpty {
	//				v = NoData
	//			}
	//			xl.Data = append(xl.Data, v)
	//			item.ResearchResult.RegisterPropertySets[i].writeXlsxAggregatedValues(v)
	//			for _, other := range item.ResearchResult.RegisterPropertySets[i].RegisterPropertyOthers {
	//				v := No
	//				for _, setValue := range g.Research.RegisterGroupsToPatients[g.PatientIndex].RegisterPropertyOthersToPatient {
	//					if setValue.RegisterPropertyOtherID == other.ID {
	//						v = setValue.Value
	//					}
	//				}
	//				if writeEmpty {
	//					v = NoData
	//				}
	//				xl.Data = append(xl.Data, v)
	//			}
	//		}
	//	}
	//	if item.ResearchResult.ValueType.IsRadio() {
	//		for i := range item.ResearchResult.AnswerVariants {
	//			exists := No
	//			for _, propToPat := range g.Research.RegisterGroupsToPatients[g.PatientIndex].RegisterPropertyToPatient {
	//				if propToPat.AnswerVariantID == item.ResearchResult.AnswerVariants[i].ID {
	//					exists = Yes
	//				}
	//			}
	//			if writeEmpty {
	//				exists = NoData
	//			}
	//			item.ResearchResult.AnswerVariants[i].writeXlsxAggregatedValues(exists)
	//			xl.Data = append(xl.Data, exists)
	//			for _, other := range item.ResearchResult.AnswerVariants[i].RegisterPropertyOthers {
	//				v := No
	//				for _, setValue := range g.Research.RegisterGroupsToPatients[g.PatientIndex].RegisterPropertyOthersToPatient {
	//					if setValue.RegisterPropertyOtherID == other.ID {
	//						v = setValue.Value
	//					}
	//				}
	//				if writeEmpty {
	//					v = NoData
	//				}
	//				xl.Data = append(xl.Data, v)
	//			}
	//		}
	//	}
	//} else {
	//	res := g.GetResultFromData(item.ResearchResult, propNum)
	//	if writeEmpty {
	//		res = NoData
	//	}
	//	xl.Data = append(xl.Data, res)
	//	item.writeXlsxAggregatedValues(res)
	//}
}

func (item *ResearchQueryGroupQuestion) writeXlsxAggregatedValues(key string) {
	_, ok := item.AggregatedValues[key]
	if ok {
		item.AggregatedValues[key]++
	} else {
		item.AggregatedValues[key] = 1
	}
}

func (items ResearchQueryGroupQuestions) writeAggregates(xl *xlsxhelper.XlsxHelper) {
	for i := range items {
		items[i].writeAggregates(xl)
	}
}

func (item *ResearchQueryGroupQuestion) writeAggregates(xl *xlsxhelper.XlsxHelper) {
	//if item.EveryRadioSet {
	//	if item.ResearchResult.ValueType.IsSet() {
	//		for _, s := range item.ResearchResult.RegisterPropertySets {
	//			item.AggregateType.WriteAggregatedValues(xl, s.AggregatedValues)
	//			for range s.RegisterPropertyOthers {
	//				xl.Cursor++
	//			}
	//		}
	//	}
	//	if item.ResearchResult.ValueType.IsRadio() {
	//		for _, r := range item.ResearchResult.AnswerVariants {
	//			item.AggregateType.WriteAggregatedValues(xl, r.AggregatedValues)
	//			for range r.RegisterPropertyOthers {
	//				xl.Cursor++
	//			}
	//		}
	//	}
	//} else {
	//	item.AggregateType.WriteAggregatedValues(xl, item.AggregatedValues)
	//}
}
