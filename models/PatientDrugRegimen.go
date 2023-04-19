package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientDrugRegimen struct {
	bun.BaseModel `bun:"patient_drug_regimens,alias:patient_drug_regimens"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Date          time.Time     `json:"date"`
	PatientID     uuid.NullUUID `bun:"type:uuid" json:"patientId"`
	DrugRegimenID uuid.UUID     `bun:"type:uuid" json:"drugRegimenId"`
	DeletedAt     *time.Time    `bun:",soft_delete" json:"deletedAt"`

	Patient     *Patient     `bun:"rel:belongs-to" json:"patient"`
	DrugRegimen *DrugRegimen `bun:"rel:belongs-to" json:"drugRegimen"`

	PatientDrugRegimenItems []*PatientDrugRegimenItem `bun:"rel:has-many" json:"patientDrugRegimenItems"`
}

func (item *PatientDrugRegimen) SetIDForChildren() {
	if len(item.PatientDrugRegimenItems) > 0 {
		for i := range item.PatientDrugRegimenItems {
			item.PatientDrugRegimenItems[i].PatientDrugRegimenID = item.ID
		}
	}
}

func GetPatientDrugRegimenItems(items []*PatientDrugRegimen) []*PatientDrugRegimenItem {
	itemsForGet := make([]*PatientDrugRegimenItem, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		items[i].SetIDForChildren()
		itemsForGet = append(itemsForGet, items[i].PatientDrugRegimenItems...)
	}
	return itemsForGet
}
