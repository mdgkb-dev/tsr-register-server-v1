package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Formula struct {
	bun.BaseModel `bun:"formulas,alias:formulas"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string    `json:"name"`
	Formula       string    `json:"formula"`

	Research   *Research `bun:"rel:belongs-to" json:"research"`
	ResearchID uuid.UUID `bun:"type:uuid" json:"researchId"`
}

type Formulas []*Formula

//
//func (item *Formula) SetIDForChildren() {
//	if len(item.AnswersVariants) > 0 {
//		for i := range item.AnswersVariants {
//			item.AnswersVariants[i].FormulaID = item.ID
//		}
//	}
//}
//
//func (items Formulas) SetIDForChildren() {
//	if len(items) == 0 {
//		return
//	}
//	for i := range items {
//		items[i].SetIDForChildren()
//	}
//}
//
//func (items Formulas) GetRegisterPropertyExamples() FormulaExamples {
//	itemsForGet := make(FormulaExamples, 0)
//	for i := range items {
//		itemsForGet = append(itemsForGet, items[i].FormulaExamples...)
//	}
//	return itemsForGet
//}
//
//func (items Formulas) GetRegisterPropertyRadios() AnswersVariants {
//	itemsForGet := make(AnswersVariants, 0)
//	for i := range items {
//		itemsForGet = append(itemsForGet, items[i].AnswersVariants...)
//	}
//	return itemsForGet
//}
//
//func (items Formulas) GetRegisterPropertyRadioForDelete() []uuid.UUID {
//	itemsForGet := make([]uuid.UUID, 0)
//	for i := range items {
//		itemsForGet = append(itemsForGet, items[i].AnswersVariantsForDelete...)
//	}
//	return itemsForGet
//}
//
//func (items Formulas) GetRegisterPropertyExamplesForDelete() []uuid.UUID {
//	itemsForGet := make([]uuid.UUID, 0)
//	for i := range items {
//		itemsForGet = append(itemsForGet, items[i].FormulaExamplesForDelete...)
//	}
//	return itemsForGet
//}
//
//func (items Formulas) GetRegisterPropertyMeasuresForDelete() []uuid.UUID {
//	itemsForGet := make([]uuid.UUID, 0)
//	for i := range items {
//		itemsForGet = append(itemsForGet, items[i].FormulaMeasuresForDelete...)
//	}
//	return itemsForGet
//}
//
//func (items Formulas) GetRegisterPropertyMeasures() FormulaMeasures {
//	itemsForGet := make(FormulaMeasures, 0)
//	for i := range items {
//		itemsForGet = append(itemsForGet, items[i].FormulaMeasures...)
//	}
//	return itemsForGet
//}
//
//func (items Formulas) GetRegisterPropertyVariants() FormulaVariants {
//	itemsForGet := make(FormulaVariants, 0)
//	for i := range items {
//		itemsForGet = append(itemsForGet, items[i].FormulaVariants...)
//	}
//	return itemsForGet
//}
//
//func (items Formulas) GetRegisterPropertyVariantsForDelete() []uuid.UUID {
//	itemsForGet := make([]uuid.UUID, 0)
//	for i := range items {
//		itemsForGet = append(itemsForGet, items[i].FormulaVariantsForDelete...)
//	}
//	return itemsForGet
//}
