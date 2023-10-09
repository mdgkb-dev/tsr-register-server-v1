package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResearchResult struct {
	bun.BaseModel `bun:"research_results,alias:research_results"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Date *time.Time `bun:"item_date" json:"date"`

	PatientResearch   *PatientResearch `bun:"rel:belongs-to" json:"patientResearch"`
	PatientResearchID uuid.NullUUID    `bun:"type:uuid" json:"patientResearchId"`
	Answers           Answers          `bun:"rel:has-many" json:"answers"`

	Research   *Research     `bun:"rel:belongs-to" json:"research"`
	ResearchID uuid.NullUUID `bun:"type:uuid" json:"researchId"`

	Patient   *Patient      `bun:"rel:belongs-to" json:"patient"`
	PatientID uuid.NullUUID `bun:"type:uuid" json:"patientId"`

	FillingPercentage uint `json:"fillingPercentage"`
	Order             uint `bun:"item_order" json:"order"`
}

type ResearchResults []*ResearchResult

type ResearchResultsWithCount struct {
	ResearchResults ResearchResults `json:"items"`
	Count           int             `json:"count"`
}

func (items ResearchResults) GetLastResult() *ResearchResult {
	var lastResult *ResearchResult
	for i := range items {
		if lastResult == nil || lastResult.Date.Before(*items[i].Date) {
			lastResult = items[i]
		}
	}
	return lastResult
}

func (items ResearchResults) SetFilePath(fileID *string) *string {
	for i := range items {
		path := items[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}

func (item *ResearchResult) SetFilePath(fileID *string) *string {
	for i := range item.Answers {
		path := item.Answers[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}

func (item *ResearchResult) SetIDForChildren() {
	//if len(item.RegisterPropertyToPatient) > 0 {
	//	for i := range item.RegisterPropertyToPatient {
	//		item.RegisterPropertyToPatient[i].ResearchResultID = item.ID
	//	}
	//}
	if len(item.Answers) > 0 {
		for i := range item.Answers {
			item.Answers[i].ResearchResultID = item.ID
		}
	}
	//if len(item.RegisterPropertyOthersToPatient) > 0 {
	//	for i := range item.RegisterPropertyOthersToPatient {
	//		item.RegisterPropertyOthersToPatient[i].ResearchResultID = item.ID
	//	}
	//}
}

func (items ResearchResults) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (items ResearchResults) SetDeleteIDForChildren() {
	//for i := range items {
	//	items[i].RegisterPropertyToPatientForDelete = append(items[i].RegisterPropertyToPatientForDelete, item.Answer[i].ID)
	//}
	//for i := range item.Answer {
	//	item.RegisterPropertySetToPatientForDelete = append(item.RegisterPropertySetToPatientForDelete, item.Answer[i].ID)
	//}
}

func (items ResearchResults) GetRegisterPropertiesToPatients() Answers {
	itemsForGet := make(Answers, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	//for i := range items {
	//	itemsForGet = append(itemsForGet, items[i].RegisterPropertyToPatient...)
	//}

	return itemsForGet
}

func (items ResearchResults) GetAnswers() Answers {
	itemsForGet := make(Answers, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	return itemsForGet
}

func (items ResearchResults) GetRegisterPropertiesToPatientsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	//for i := range items {
	//	itemsForGet = append(itemsForGet, items[i].RegisterPropertyToPatientForDelete...)
	//}
	return itemsForGet
}

func (items ResearchResults) GetRegisterPropertySetToPatient() []*Answer {
	itemsForGet := make([]*Answer, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	//for i := range items {
	//	itemsForGet = append(itemsForGet, items[i].Answer...)
	//}
	return itemsForGet
}

func (items ResearchResults) GetRegisterPropertySetToPatientForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	//for i := range items {
	//	itemsForGet = append(itemsForGet, items[i].RegisterPropertySetToPatientForDelete...)
	//}
	return itemsForGet
}

func (items ResearchResults) GetRegisterPropertyOthersToPatient() PatientAnswerComments {
	itemsForGet := make(PatientAnswerComments, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	//for i := range items {
	//	itemsForGet = append(itemsForGet, items[i].RegisterPropertyOthersToPatient...)
	//}
	return itemsForGet
}

func (item *ResearchResult) GetAggregateExistingData() string {
	res := No
	for _, answer := range item.Answers {
		if answer.GetAggregateExistingData() {
			res = Yes
			break
		}
	}
	return res
}

func (item *ResearchResult) Include(variantID uuid.NullUUID) string {
	res := No
	for _, answer := range item.Answers {
		res = answer.AnswerVariantSelected(variantID)
		if res == Yes {
			break
		}
	}
	return res
}

func (item *ResearchResult) GetData(question *Question) interface{} {
	var res interface{}
	res = No
	for _, answer := range item.Answers {
		if answer.QuestionID == question.ID {
			res = answer.GetData(question)
			break
		}
	}
	return res
}

func (item *ResearchResult) GetScores(q *Question) int {
	sumScores := 0
	for _, answer := range item.Answers {
		if answer.QuestionID == q.ID {
			for _, radio := range q.AnswerVariants {
				if radio.ID == answer.AnswerVariantID {
					fmt.Println(radio.Score)
					sumScores += radio.Score
					break
				}
			}
		}
	}
	return sumScores
}

func (items ResearchResults) GetExportData(research *Research) ([][]interface{}, error) {
	results := make([][]interface{}, 0)
	for _, researchResult := range items {
		result, err := researchResult.GetXlsxData(research)
		if err != nil {
			break
		}
		results = append(results, result)
	}
	return results, nil
}

func (item *ResearchResult) GetXlsxData(research *Research) ([]interface{}, error) {
	result := make([]interface{}, 0)
	result = append(result, item.Date)
	//if research.WithScores {
	//	sum := 0
	//	for _, q := range research.Questions {
	//		sum += item.GetScores(q)
	//	}
	//	results[resultN] = append(results[resultN], strconv.Itoa(sum))
	//	return err
	//}

	variables := make(map[string]interface{})
	for _, q := range research.Questions {
		answer := item.GetData(q)
		result = append(result, answer)
		variables[q.Code] = answer
	}
	resultFormulas, err := research.Formulas.SetXlsxData(variables)
	if err != nil {
		return nil, err
	}
	result = append(result, resultFormulas...)
	return result, nil
}

func (item *ResearchResult) GetResultsMap(questions Questions) map[string]interface{} {
	variables := make(map[string]interface{})
	for _, q := range questions {
		answer := item.GetData(q)
		variables[q.Code] = answer
	}
	return variables
}

func (item *ResearchResult) GetAnthropometry() (uint, uint) {
	height, weight := 0, 0
	for _, answer := range item.Answers {
		if answer.Question.Code == string(AnthropomethryKeyHeight) {
			height = int(answer.ValueNumber)
		}
		if answer.Question.Code == string(AnthropomethryKeyWeight) {
			weight = int(answer.ValueNumber)
		}
	}
	return uint(height), uint(weight)
}
