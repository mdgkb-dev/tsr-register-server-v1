package models

import (
	"sort"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type AnswerVariant struct {
	bun.BaseModel `bun:"answer_variants,alias:answer_variants"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	QuestionID    uuid.NullUUID `bun:"type:uuid" json:"questionId"`
	Question      *Question     `bun:"rel:belongs-to" json:"question"`
	Order         int           `bun:"item_order" json:"order"`
	Score         int           `json:"score"`
	//RegisterPropertyOthers          RegisterPropertyOthers `bun:"rel:has-many" json:"registerPropertyOthers"`
	//RegisterPropertyOthersForDelete []uuid.UUID            `bun:"-" json:"registerPropertyOthersForDelete"`

	AggregatedValues         map[string]float64       `bun:"-" json:"aggregatedValues"`
	RegisterQueryPercentages RegisterQueryPercentages `bun:"-" `
}

type AnswerVariants []*AnswerVariant

func (item *AnswerVariant) SetIDForChildren() {
	//if len(item.RegisterPropertyOthers) == 0 {
	//	return
	//}
	//for i := range item.RegisterPropertyOthers {
	//	item.RegisterPropertyOthers[i].AnswerVariantID = item.ID
	//}
}

func (items AnswerVariants) SetIDForChildren() {
	if len(items) == 0 {
		return
	}
	for i := range items {
		items[i].SetIDForChildren()
	}
}

//func (items AnswerVariants) GetRegisterPropertyOthers() RegisterPropertyOthers {
//	itemsForGet := make(RegisterPropertyOthers, 0)
//	for i := range items {
//		itemsForGet = append(itemsForGet, items[i].RegisterPropertyOthers...)
//	}
//	return itemsForGet
//}

func (items AnswerVariants) GetRegisterPropertyOthersForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	//for i := range items {
	//	itemsForGet = append(itemsForGet, items[i].RegisterPropertyOthersForDelete...)
	//}
	return itemsForGet
}

func (item *AnswerVariant) writeXlsxAggregatedValues(key string) {
	_, ok := item.AggregatedValues[key]
	if ok {
		item.AggregatedValues[key]++
	} else {
		item.AggregatedValues[key] = 1
	}
}

func (item *AnswerVariant) GetAggregatedPercentage() {
	sum := float64(0)
	for k, v := range item.AggregatedValues {
		sum += v
		item.RegisterQueryPercentages = append(item.RegisterQueryPercentages, &RegisterQueryPercentage{k, v})
	}
	sort.Slice(item.RegisterQueryPercentages, func(i, j int) bool {
		return item.RegisterQueryPercentages[i].Value > item.RegisterQueryPercentages[j].Value
	})
}
