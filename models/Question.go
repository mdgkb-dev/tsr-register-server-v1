package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Question struct {
	bun.BaseModel   `bun:"questions,alias:questions"`
	ID              uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name            string        `json:"name"`
	ShortName       string        `json:"shortName"`
	Code            string        `json:"code"`
	WithOther       bool          `json:"withOther"`
	Order           int           `bun:"item_order" json:"order"`
	ValueType       *ValueType    `bun:"rel:belongs-to" json:"valueType"`
	ValueTypeID     uuid.NullUUID `bun:"type:uuid" json:"valueTypeId"`
	Research        *Research     `bun:"rel:belongs-to" json:"research"`
	ResearchID      uuid.NullUUID `bun:"type:uuid" json:"researchId"`
	AgeCompare      bool          `json:"ageCompare"`
	CalculateScores bool          `json:"calculateScores"`

	AnswerVariants          AnswerVariants `bun:"rel:has-many" json:"answerVariants"`
	AnswerVariantsForDelete []uuid.UUID    `bun:"-" json:"answerVariantsForDelete"`

	QuestionVariants          QuestionVariants `bun:"rel:has-many" json:"questionVariants"`
	QuestionVariantsForDelete []uuid.UUID      `bun:"-" json:"questionVariantsForDelete"`

	QuestionExamples          QuestionExamples `bun:"rel:has-many" json:"questionExamples"`
	QuestionExamplesForDelete []uuid.UUID      `bun:"-" json:"questionExamplesForDelete"`

	QuestionMeasures          QuestionMeasures `bun:"rel:has-many" json:"questionMeasures"`
	QuestionMeasuresForDelete []uuid.UUID      `bun:"-" json:"questionMeasuresForDelete"`
	Children                  Questions        `bun:"rel:has-many,join:id=parent_id" json:"children"`
	ParentID                  uuid.NullUUID    `bun:"type:uuid" json:"parentId"`
	Parent                    *Question        `bun:"-" json:"parent"`
}

type Questions []*Question

type QuestionsWithCount struct {
	Questions Questions `json:"items"`
	Count     int       `json:"count"`
}

func (item *Question) SetIDForChildren() {
	if len(item.AnswerVariants) > 0 {
		for i := range item.AnswerVariants {
			item.AnswerVariants[i].QuestionID = item.ID
		}
	}
	if len(item.QuestionExamples) > 0 {
		for i := range item.QuestionExamples {
			item.QuestionExamples[i].QuestionID = item.ID
		}
	}
	if len(item.QuestionMeasures) > 0 {
		for i := range item.QuestionMeasures {
			item.QuestionMeasures[i].QuestionID = item.ID
		}
	}
	if len(item.QuestionVariants) > 0 {
		for i := range item.QuestionVariants {
			item.QuestionVariants[i].QuestionID = item.ID
		}
	}
}

func (items Questions) SetIDForChildren() {
	if len(items) == 0 {
		return
	}
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (items Questions) GetRegisterPropertyExamples() QuestionExamples {
	itemsForGet := make(QuestionExamples, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].QuestionExamples...)
	}
	return itemsForGet
}

func (items Questions) GetRegisterPropertyRadios() AnswerVariants {
	itemsForGet := make(AnswerVariants, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].AnswerVariants...)
	}
	return itemsForGet
}

func (items Questions) GetRegisterPropertyRadioForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].AnswerVariantsForDelete...)
	}
	return itemsForGet
}

func (items Questions) GetRegisterPropertyExamplesForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].QuestionExamplesForDelete...)
	}
	return itemsForGet
}

func (items Questions) GetRegisterPropertyMeasuresForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].QuestionMeasuresForDelete...)
	}
	return itemsForGet
}

func (items Questions) GetRegisterPropertyMeasures() QuestionMeasures {
	itemsForGet := make(QuestionMeasures, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].QuestionMeasures...)
	}
	return itemsForGet
}

func (items Questions) GetRegisterPropertyVariants() QuestionVariants {
	itemsForGet := make(QuestionVariants, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].QuestionVariants...)
	}
	return itemsForGet
}

func (items Questions) GetRegisterPropertyVariantsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].QuestionVariantsForDelete...)
	}
	return itemsForGet
}
