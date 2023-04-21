package models

import (
	"mdgkb/tsr-tegister-server-v1/helpers/xlsxhelper"
	"sort"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterQueryGroupProperty struct {
	bun.BaseModel        `bun:"register_query_group_properties,alias:register_query_group_properties"`
	ID                   uuid.UUID           `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	RegisterQueryGroupID uuid.UUID           `bun:"type:uuid" json:"registerQueryGroupId"`
	RegisterQueryGroup   *RegisterQueryGroup `bun:"rel:belongs-to" json:"registerQueryGroup"`

	RegisterPropertyID uuid.UUID `bun:"type:uuid" json:"registerPropertyId"`
	RegisterProperty   *Question `bun:"rel:belongs-to" json:"registerProperty"`

	Order                    int                      `bun:"item_order" json:"order"`
	AggregatedValues         map[string]float64       `bun:"-" json:"aggregatedValues"`
	RegisterQueryPercentages RegisterQueryPercentages `bun:"-" `
	EveryRadioSet            bool                     `json:"everyRadioSet"`

	AggregateType AggregateType `json:"aggregateType"`
}

type RegisterQueryGroupProperties []*RegisterQueryGroupProperty

func (item *RegisterQueryGroupProperty) GetAggregatedPercentage() {
	sum := float64(0)
	for k, v := range item.AggregatedValues {
		sum += v
		item.RegisterQueryPercentages = append(item.RegisterQueryPercentages, &RegisterQueryPercentage{k, v})
	}
	sort.Slice(item.RegisterQueryPercentages, func(i, j int) bool {
		return item.RegisterQueryPercentages[i].Value > item.RegisterQueryPercentages[j].Value
	})
}

func (items RegisterQueryGroupProperties) writeXlsxHeader(xl *xlsxhelper.XlsxHelper) {
	for i := range items {
		items[i].writeXlsxHeader(xl)
	}
}

func (item *RegisterQueryGroupProperty) writeXlsxHeader(xl *xlsxhelper.XlsxHelper) {
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

func (items RegisterQueryGroupProperties) writeXlsxData(xl *xlsxhelper.XlsxHelper, g *RegisterQueryGroup, writeEmpty bool) {
	for i := range items {
		items[i].writeXlsxData(xl, g, i, writeEmpty)
	}
}

func (item *RegisterQueryGroupProperty) writeXlsxData(xl *xlsxhelper.XlsxHelper, g *RegisterQueryGroup, propNum int, writeEmpty bool) {
	//if item.EveryRadioSet {
	//	if item.ResearchResult.ValueType.IsSet() {
	//		for i := range item.ResearchResult.RegisterPropertySets {
	//			v := No
	//			if g.RegisterGroup.RegisterGroupsToPatients[g.PatientIndex].RegisterPropertySetToPatient.Include(item.ResearchResult.RegisterPropertySets[i].ID) {
	//				v = Yes
	//			}
	//			if writeEmpty {
	//				v = NoData
	//			}
	//			xl.Data = append(xl.Data, v)
	//			item.ResearchResult.RegisterPropertySets[i].writeXlsxAggregatedValues(v)
	//			for _, other := range item.ResearchResult.RegisterPropertySets[i].RegisterPropertyOthers {
	//				v := No
	//				for _, setValue := range g.RegisterGroup.RegisterGroupsToPatients[g.PatientIndex].RegisterPropertyOthersToPatient {
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
	//			for _, propToPat := range g.RegisterGroup.RegisterGroupsToPatients[g.PatientIndex].RegisterPropertyToPatient {
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
	//				for _, setValue := range g.RegisterGroup.RegisterGroupsToPatients[g.PatientIndex].RegisterPropertyOthersToPatient {
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

func (item *RegisterQueryGroupProperty) writeXlsxAggregatedValues(key string) {
	_, ok := item.AggregatedValues[key]
	if ok {
		item.AggregatedValues[key]++
	} else {
		item.AggregatedValues[key] = 1
	}
}

func (items RegisterQueryGroupProperties) writeAggregates(xl *xlsxhelper.XlsxHelper) {
	for i := range items {
		items[i].writeAggregates(xl)
	}
}

func (item *RegisterQueryGroupProperty) writeAggregates(xl *xlsxhelper.XlsxHelper) {
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
