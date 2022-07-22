package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type RegisterGroupToPatient struct {
	bun.BaseModel `bun:"register_groups_to_patients,alias:register_groups_to_patients"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Date *time.Time `bun:"register_groups_to_patients_date" json:"date"`

	RegisterGroup   *RegisterGroup `bun:"rel:belongs-to" json:"registerGroup"`
	RegisterGroupID uuid.UUID      `bun:"type:uuid" json:"registerGroupId"`
	Patient         *Patient       `bun:"rel:has-one" json:"patients"`
	PatientID       uuid.UUID      `bun:"type:uuid" json:"patientId"`

	RegisterPropertyToPatient             []*RegisterPropertyToPatient    `bun:"rel:has-many" json:"registerPropertyToPatient"`
	RegisterPropertyOthersToPatient       RegisterPropertyOthersToPatient `bun:"rel:has-many" json:"registerPropertyOthersToPatient"`
	RegisterPropertyToPatientForDelete    []uuid.UUID                     `bun:"-" json:"registerPropertyToPatientForDelete"`
	RegisterPropertySetToPatient          []*RegisterPropertySetToPatient `bun:"rel:has-many" json:"registerPropertySetToPatient"`
	RegisterPropertySetToPatientForDelete []uuid.UUID                     `bun:"-" json:"registerPropertySetToPatientForDelete"`
}

type RegisterGroupsToPatients []*RegisterGroupToPatient

func (item *RegisterGroupToPatient) SetIdForChildren() {
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

func (items RegisterGroupsToPatients) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (items RegisterGroupsToPatients) SetDeleteIdForChildren() {
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
