package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResearchesPoolResearch struct {
	bun.BaseModel    `bun:"researches_pools_researches,alias:researches_pools_researches"`
	ID               uuid.UUID       `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	ResearchesPool   *ResearchesPool `bun:"rel:belongs-to" json:"researchesPool"`
	ResearchesPoolID uuid.UUID       `bun:"type:uuid" json:"researchesPoolId"`

	Research   *Research     `bun:"rel:belongs-to" json:"research"`
	ResearchID uuid.NullUUID `bun:"type:uuid" json:"researchId"`

	Order uint `bun:"item_order" json:"order"`
}

type ResearchesPoolsResearches []*ResearchesPoolResearch

//func (item *Research) SetIDForChildren() {
//	if len(item.Questions) == 0 {
//		return
//	}
//	for i := range item.Questions {
//		item.Questions[i].ResearchID = item.ID
//	}
//}
//
//func (items Researches) SetIDForChildren() {
//	if len(items) == 0 {
//		return
//	}
//	for i := range items {
//		items[i].SetIDForChildren()
//	}
//}
//
//func (items Researches) GetQuestions() Questions {
//	itemsForGet := make(Questions, 0)
//	for i := range items {
//		itemsForGet = append(itemsForGet, items[i].Questions...)
//	}
//	return itemsForGet
//}
//
//func (items Researches) GetQuestionsForDelete() []uuid.UUID {
//	itemsForGet := make([]uuid.UUID, 0)
//	for i := range items {
//		itemsForGet = append(itemsForGet, items[i].QuestionsForDelete...)
//	}
//	return itemsForGet
//}
