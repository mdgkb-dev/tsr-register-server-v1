package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterGroupToPatient struct {
	bun.BaseModel `bun:"register_groups_to_patients,alias:register_groups_to_patients"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Date *time.Time `bun:"register_groups_to_patients_date" json:"date"`

	RegisterGroup                         *RegisterGroup                  `bun:"rel:belongs-to" json:"registerGroup"`
	RegisterGroupID                       uuid.UUID                       `bun:"type:uuid" json:"registerGroupId"`
	Patient                               *Patient                        `bun:"rel:belongs-to" json:"patients"`
	PatientID                             uuid.UUID                       `bun:"type:uuid" json:"patientId"`
	RegisterPropertyToPatient             []*RegisterPropertyToPatient    `bun:"rel:has-many" json:"registerPropertyToPatient"`
	RegisterPropertyOthersToPatient       RegisterPropertyOthersToPatient `bun:"rel:has-many" json:"registerPropertyOthersToPatient"`
	RegisterPropertyToPatientForDelete    []uuid.UUID                     `bun:"-" json:"registerPropertyToPatientForDelete"`
	RegisterPropertySetToPatient          RegisterPropertySetsToPatients  `bun:"rel:has-many" json:"registerPropertySetToPatient"`
	RegisterPropertySetToPatientForDelete []uuid.UUID                     `bun:"-" json:"registerPropertySetToPatientForDelete"`
}

type RegisterGroupsToPatients []*RegisterGroupToPatient

func (items RegisterGroupsToPatients) SetFilePath(fileID *string) *string {
	for i := range items {
		path := items[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}

func (item *RegisterGroupToPatient) SetFilePath(fileID *string) *string {
	for i := range item.RegisterPropertyToPatient {
		path := item.RegisterPropertyToPatient[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}

func (item *RegisterGroupToPatient) SetIDForChildren() {
	if len(item.RegisterPropertyToPatient) > 0 {
		for i := range item.RegisterPropertyToPatient {
			item.RegisterPropertyToPatient[i].RegisterGroupToPatientID = item.ID
		}
	}
	if len(item.RegisterPropertySetToPatient) > 0 {
		for i := range item.RegisterPropertySetToPatient {
			item.RegisterPropertySetToPatient[i].RegisterGroupToPatientID = item.ID
		}
	}
	if len(item.RegisterPropertyOthersToPatient) > 0 {
		for i := range item.RegisterPropertyOthersToPatient {
			item.RegisterPropertyOthersToPatient[i].RegisterGroupToPatientID = item.ID
		}
	}
}

func (items RegisterGroupsToPatients) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (items RegisterGroupsToPatients) SetDeleteIDForChildren() {
	//for i := range items {
	//	items[i].RegisterPropertyToPatientForDelete = append(items[i].RegisterPropertyToPatientForDelete, item.RegisterPropertyToPatient[i].ID)
	//}
	//for i := range item.RegisterPropertySetToPatient {
	//	item.RegisterPropertySetToPatientForDelete = append(item.RegisterPropertySetToPatientForDelete, item.RegisterPropertySetToPatient[i].ID)
	//}
}

func (items RegisterGroupsToPatients) GetRegisterPropertiesToPatients() RegisterPropertiesToPatients {
	itemsForGet := make(RegisterPropertiesToPatients, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyToPatient...)
	}

	return itemsForGet
}

func (items RegisterGroupsToPatients) GetRegisterPropertiesToPatientsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyToPatientForDelete...)
	}
	return itemsForGet
}

func (items RegisterGroupsToPatients) GetRegisterPropertySetToPatient() []*RegisterPropertySetToPatient {
	itemsForGet := make([]*RegisterPropertySetToPatient, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertySetToPatient...)
	}
	return itemsForGet
}

func (items RegisterGroupsToPatients) GetRegisterPropertySetToPatientForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertySetToPatientForDelete...)
	}
	return itemsForGet
}

func (items RegisterGroupsToPatients) GetRegisterPropertyOthersToPatient() RegisterPropertyOthersToPatient {
	itemsForGet := make(RegisterPropertyOthersToPatient, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyOthersToPatient...)
	}
	return itemsForGet
}

//func (items RegisterGroupsToPatients) GetRegisterPropertySetToPatientForDelete() []uuid.UUID {
//	itemsForGet := make([]uuid.UUID, 0)
//	if len(items) == 0 {
//		return itemsForGet
//	}
//	for i := range items {
//		itemsForGet = append(itemsForGet, items[i].RegisterPropertySetToPatientForDelete...)
//	}
//	return itemsForGet
//}

func (item *RegisterGroupToPatient) GetAggregateExistingData() string {
	res := No
	for _, prop := range item.RegisterPropertyToPatient {
		if prop.GetAggregateExistingData() {
			res = Yes
			break
		}
	}
	return res
}
