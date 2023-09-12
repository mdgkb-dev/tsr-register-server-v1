package models

import (
	"encoding/json"
	"errors"

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

func (items Researches) GetExportData() [][]interface{} {
	headersLines := make([][]interface{}, 0)
	headersLines = append(headersLines, []interface{}{})
	for _, research := range items {
		researchHeaders := research.GetHeaders()
		for _, researchHeader := range researchHeaders {
			headersLines[0] = append(headersLines[0], researchHeader...)
		}
	}
	return headersLines
}

func (item *Research) GetHeaders() [][]interface{} {
	headersLines := make([][]interface{}, 0)
	headersLines = append(headersLines, []interface{}{})
	headersLines[0] = append(headersLines[0], "Дата")

	if item.WithScores {
		headersLines[0] = append(headersLines[0], "Всего баллов")
		headersLines[0] = append(headersLines[0], "Всего баллов по шкале")
		return headersLines
	}

	for _, q := range item.Questions {
		headersLines[0] = append(headersLines[0], q.Name)
	}
	for _, f := range item.Formulas {
		if f.Xlsx {
			headersLines[0] = append(headersLines[0], f.Name)
		}
		if len(f.FormulaResults) > 0 {
			headersLines[0] = append(headersLines[0], "Результат")
		}
	}

	return headersLines
}

type ResearchesExport struct {
	IDPool          []string `json:"ids"`
	WithAge         bool     `json:"withAge"`
	CountAverageAge bool     `json:"countAverageAge"`
}

const researchesExportOptionsKey = "research"

func (item *ResearchesExport) ParseExportOptions(options map[string]map[string]interface{}) error {
	opt, ok := options[researchesExportOptionsKey]
	if !ok {
		return errors.New("not find patients")
	}
	jsonbody, err := json.Marshal(opt[researchesExportOptionsKey])
	if err != nil {
		return errors.New("parse error")
	}

	if err := json.Unmarshal(jsonbody, &item); err != nil {
		return errors.New("parse error")
	}
	return nil
}
