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
		xl.WriteString(2, xl.Cursor, &[]string{item.Question.Name})
		if item.Question.ValueType.IsSet() {
			for _, setItem := range item.Question.AnswerVariants {
				xl.WriteString(3, xl.Cursor, &[]string{setItem.Name})
				xl.Cursor++
			}
		}
		if item.Question.ValueType.IsRadio() {
			for _, radioItem := range item.Question.AnswerVariants {
				xl.WriteString(3, xl.Cursor, &[]string{radioItem.Name})
				xl.Cursor++
			}
		}
		for _, childQuestion := range item.Question.Children {
			xl.WriteString(3, xl.Cursor, &[]string{childQuestion.Name})
			xl.Cursor++
		}
	} else {
		xl.WriteString(2, xl.Cursor, &[]string{item.Question.Name})
		xl.Cursor++
	}
}

func (items ResearchQueryGroupQuestions) writeXlsxData(xl *xlsxhelper.XlsxHelper, g *ResearchQueryGroup, result *ResearchResult) {
	for i := range items {
		items[i].writeXlsxData(xl, g, i, result)
	}
}

func (item *ResearchQueryGroupQuestion) writeXlsxData(xl *xlsxhelper.XlsxHelper, g *ResearchQueryGroup, propNum int, result *ResearchResult) {
	if item.EveryRadioSet {
		if item.Question.ValueType.IsSet() {

			for _, av := range item.Question.AnswerVariants {
				v := result.Include(av.ID)
				if item.Question.Name == "Если у Вас при развитии данной реакции отмечались симптомы со стороны кожи/слизистых, то какие?" && g.PatientIndex == 0 {
					fmt.Println(v)
				}
				xl.Data = append(xl.Data, v)
				//item.ResearchResult.RegisterPropertySets[i].writeXlsxAggregatedValues(v)
			}
		}
		if item.Question.ValueType.IsRadio() {
			for i := range item.Question.AnswerVariants {
				exists := No
				for _, answer := range result.Answers {
					if answer.AnswerVariantID == item.Question.AnswerVariants[i].ID {
						exists = Yes
					}
				}
				xl.Data = append(xl.Data, exists)
			}
			//item.writeXlsxAggregatedValues(exists)
		}
		for _, childQuestion := range item.Question.Children {
			a := result.GetData(childQuestion)
			xl.Data = append(xl.Data, a)
		}
	} else {
		res := g.GetResultFromData(item.Question, result)
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
