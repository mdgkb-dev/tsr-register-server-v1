package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Research struct {
	bun.BaseModel `bun:"researches,alias:researches"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`

	Questions          Questions   `bun:"rel:has-many" json:"questions"`
	QuestionsForDelete []uuid.UUID `bun:"-" json:"questionsForDelete"`
	WithDates          bool        `json:"withDates"`
	WithScores         bool        `json:"withScores"`

	ResearchResults ResearchResults `bun:"rel:has-many" json:"researchResults"`

	Formulas          Formulas    `bun:"rel:has-many" json:"formulas"`
	FormulasForDelete []uuid.UUID `bun:"-" json:"formulasForDelete"`
}

type Researches []*Research

func (item *Research) SetIDForChildren() {
	if len(item.Questions) == 0 {
		return
	}
	for i := range item.Questions {
		item.Questions[i].ResearchID = item.ID
	}
}

func (items Researches) SetIDForChildren() {
	if len(items) == 0 {
		return
	}
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (items Researches) GetQuestions() Questions {
	itemsForGet := make(Questions, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].Questions...)
	}
	return itemsForGet
}

func (items Researches) GetQuestionsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].QuestionsForDelete...)
	}
	return itemsForGet
}

func (item *Research) GetResultByPatientID(patientID uuid.NullUUID) *ResearchResult {
	research := &ResearchResult{}
	for i := range item.ResearchResults {
		if item.ResearchResults[i].PatientID == patientID {
			research = item.ResearchResults[i]
		}
	}
	return research
}
