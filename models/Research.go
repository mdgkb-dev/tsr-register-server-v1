package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Research struct {
	bun.BaseModel      `bun:"researches,alias:researches"`
	ID                 uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name               string        `json:"name"`
	IsLaboratory       bool          `json:"isLaboratory"`
	Questions          Questions     `bun:"rel:has-many" json:"questions"`
	QuestionsForDelete []uuid.UUID   `bun:"-" json:"questionsForDelete"`
	WithDates          bool          `json:"withDates"`
	WithScores         bool          `json:"withScores"`

	ResearchResults ResearchResults `bun:"rel:has-many" json:"researchResults"`

	Formulas          Formulas    `bun:"rel:has-many" json:"formulas"`
	FormulasForDelete []uuid.UUID `bun:"-" json:"formulasForDelete"`

	Order uint `bun:"item_order" json:"order"`
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

func (item *Research) GetHeaders(patientName string) [][]string {

	headersLines := make([][]string, 0)
	headersLines = append(headersLines, []string{item.Name + ": " + patientName})

	headersLines = append(headersLines, []string{})
	headersLines[1] = append(headersLines[1], "Дата")

	if item.WithScores {
		headersLines[1] = append(headersLines[1], "Всего баллов")
		headersLines[1] = append(headersLines[1], "Всего баллов по шкале")
		return headersLines
	}

	for _, q := range item.Questions {
		headersLines[1] = append(headersLines[1], q.Name)
	}
	for _, f := range item.Formulas {
		if f.Xlsx {
			headersLines[1] = append(headersLines[1], f.Name)
		}
	}

	return headersLines
}
