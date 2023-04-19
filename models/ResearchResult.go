package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResearchResult struct {
	bun.BaseModel `bun:"research_results,alias:research_results"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Date *time.Time `bun:"item_date" json:"date"`

	Research   *Research     `bun:"rel:belongs-to" json:"research"`
	ResearchID uuid.UUID     `bun:"type:uuid" json:"researchId"`
	Patient    *Patient      `bun:"rel:belongs-to" json:"patients"`
	PatientID  uuid.NullUUID `bun:"type:uuid" json:"patientId"`
	Answers    Answers       `bun:"rel:has-many" json:"answers"`

	FillingPercentage uint `json:"fillingPercentage"`
	Order             uint `bun:"item_order" json:"order"`
	//PatientAnswerComments PatientAnswerComments `bun:"rel:has-many" json:"patientAnswerComments"`
	//RegisterPropertyToPatientForDelete    []uuid.UUID                     `bun:"-" json:"registerPropertyToPatientForDelete"`
	//Answer          AnswersVariants  `bun:"rel:has-many" json:"registerPropertySetToPatient"`
	//RegisterPropertySetToPatientForDelete []uuid.UUID                     `bun:"-" json:"registerPropertySetToPatientForDelete"`
}

type ResearchResults []*ResearchResult

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
	//if len(item.Answer) > 0 {
	//	for i := range item.Answer {
	//		item.Answer[i].ResearchResultID = item.ID
	//	}
	//}
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
