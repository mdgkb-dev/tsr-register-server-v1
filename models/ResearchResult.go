package models

import (
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

	FillingPercentage uint `json:"fillingPercentage"`
	Order             uint `bun:"item_order" json:"order"`
}

type ResearchResults []*ResearchResult

type ResearchResultsWithCount struct {
	ResearchResults ResearchResults `json:"items"`
	Count           int             `json:"count"`
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
	//for i := range item.RegisterPropertyToPatient {
	//	path := item.RegisterPropertyToPatient[i].SetFilePath(fileID)
	//	if path != nil {
	//		return path
	//	}
	//}
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

//func (items ResearchResults) GetRegisterPropertySetToPatientForDelete() []uuid.UUID {
//	itemsForGet := make([]uuid.UUID, 0)
//	if len(items) == 0 {
//		return itemsForGet
//	}
//	for i := range items {
//		itemsForGet = append(itemsForGet, items[i].RegisterPropertySetToPatientForDelete...)
//	}
//	return itemsForGet
//}
//
//func (item *ResearchResult) GetAggregateExistingData() string {
//	res := No
//	for _, prop := range item.RegisterPropertyToPatient {
//		if prop.GetAggregateExistingData() {
//			res = Yes
//			break
//		}
//	}
//	return res
//}
