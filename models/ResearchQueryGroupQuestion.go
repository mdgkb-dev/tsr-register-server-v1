package models

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/helpers/xlsxhelper"
	"sort"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResearchQueryGroupQuestion struct {
	bun.BaseModel        `bun:"research_query_group_questions,alias:research_query_group_questions"`
	ID                   uuid.UUID           `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	ResearchQueryGroupID uuid.NullUUID       `bun:"type:uuid" json:"researchQueryGroupId"`
	ResearchQueryGroup   *ResearchQueryGroup `bun:"rel:belongs-to" json:"researchQueryGroup"`

	//ResearchResultID uuid.UUID       `bun:"type:uuid" json:"researchResultId"`
	ResearchResult *ResearchResult `bun:"rel:belongs-to" json:"researchResult"`

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
	if item.EveryRadioSet {
		//xl.WriteString(2, xl.Cursor, &[]string{item.ResearchResult.Name})
		//if item.ResearchResult.ValueType.IsSet() {
		//	for _, setItem := range item.ResearchResult.RegisterPropertySets {
		//		xl.WriteString(3, xl.Cursor, &[]string{setItem.Name})
		//		xl.Cursor++
		//		for _, other := range setItem.RegisterPropertyOthers {
		//			xl.WriteString(3, xl.Cursor, &[]string{other.Name})
		//			xl.Cursor++
		//		}
		//	}
		//}
		//if item.ResearchResult.ValueType.IsRadio() {
		//	for _, radioItem := range item.ResearchResult.AnswerVariants {
		//		xl.WriteString(3, xl.Cursor, &[]string{radioItem.Name})
		//		xl.Cursor++
		//		for _, other := range radioItem.RegisterPropertyOthers {
		//			xl.WriteString(3, xl.Cursor, &[]string{other.Name})
		//			xl.Cursor++
		//		}
		//	}
		//}
	} else {
		xl.WriteString(2, xl.Cursor, &[]string{item.Question.Name})
		xl.Cursor++
	}
}

func (items ResearchQueryGroupQuestions) writeXlsxData(xl *xlsxhelper.XlsxHelper, g *ResearchQueryGroup, writeEmpty bool) {
	fmt.Println(len(items))
	for i := range items {
		items[i].writeXlsxData(xl, g, i, writeEmpty)
	}
}

func (item *ResearchQueryGroupQuestion) writeXlsxData(xl *xlsxhelper.XlsxHelper, g *ResearchQueryGroup, propNum int, writeEmpty bool) {
	if item.EveryRadioSet {
		//if item.Question.ValueType.IsSet() {
		//for i := range item.ResearchResult.Answers {
		v := No
		//if g.Research.ResearchResults[g.PatientIndex].Include(item.ResearchResult.RegisterPropertySets[i].ID) {
		//	v = Yes
		//}
		//if writeEmpty {
		//	v = NoData
		//}
		xl.Data = append(xl.Data, v)
		//item.ResearchResult.RegisterPropertySets[i].writeXlsxAggregatedValues(v)
		//for _, other := range item.ResearchResult.RegisterPropertySets[i].RegisterPropertyOthers {
		//	v := No
		//	for _, setValue := range g.Research.RegisterGroupsToPatients[g.PatientIndex].RegisterPropertyOthersToPatient {
		//		if setValue.RegisterPropertyOtherID == other.ID {
		//			v = setValue.Value
		//		}
		//	}
		//	if writeEmpty {
		//		v = NoData
		//	}
		//	xl.Data = append(xl.Data, v)
		//}
		//}
		//}
		if item.Question.ValueType.IsRadio() {
			exists := No
			for i := range item.Question.AnswerVariants {

				for _, answer := range g.Research.ResearchResults[g.PatientIndex].Answers {
					if answer.AnswerVariantID == item.Question.AnswerVariants[i].ID {
						exists = Yes
					}
				}

				if writeEmpty {
					exists = NoData
				}

				fmt.Println(len(xl.Data))
				//for _, other := range item.ResearchResult.AnswerVariants[i].RegisterPropertyOthers {
				//	v := No
				//	for _, setValue := range g.Research.RegisterGroupsToPatients[g.PatientIndex].RegisterPropertyOthersToPatient {
				//		if setValue.RegisterPropertyOtherID == other.ID {
				//			v = setValue.Value
				//		}
				//	}
				//	if writeEmpty {
				//		v = NoData
				//	}
				//	xl.Data = append(xl.Data, v)
				//}
			}
			xl.Data = append(xl.Data, exists)
			item.writeXlsxAggregatedValues(exists)
		}
	} else {
		res := g.GetResultFromData(item.Question, propNum)
		if writeEmpty {
			res = NoData
		}
		xl.Data = append(xl.Data, res)
		item.writeXlsxAggregatedValues(res)
	}
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
	if item.Question.ValueType.IsRadio() {
		//for _, r := range item.ResearchResult.Answers {
		item.AggregateType.WriteAggregatedValues(xl, item.AggregatedValues)
		//for range r.RegisterPropertyOthers {
		//	xl.Cursor++
		//}
		//}
	}
	//} else {
	//	item.AggregateType.WriteAggregatedValues(xl, item.AggregatedValues)
	//}
}
