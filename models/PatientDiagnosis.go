package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientDiagnosis struct {
	bun.BaseModel     `bun:"patient_diagnosis,alias:patient_diagnosis"`
	ID                uuid.UUID        `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Patient           *Patient         `bun:"rel:has-one" json:"patients"`
	PatientID         uuid.UUID        `bun:"type:uuid" json:"PatientId"`
	MkbDiagnosis      *MkbDiagnosis    `bun:"rel:belongs-to" json:"mkbDiagnosis"`
	MkbDiagnosisID    uuid.UUID        `bun:"type:uuid" json:"mkbDiagnosisId"`
	MkbSubDiagnosis   *MkbSubDiagnosis `bun:"rel:belongs-to" json:"mkbSubDiagnosis"`
	MkbSubDiagnosisID uuid.NullUUID    `bun:"type:uuid,nullzero" json:"mkbSubDiagnosisId"`
	Primary           bool             `json:"primary"`
	DeletedAt         time.Time        `bun:",soft_delete" json:"deletedAt"`

	PatientDiagnosisAnamnesis          []*PatientDiagnosisAnamnesis `bun:"rel:has-many" json:"patientDiagnosisAnamnesis"`
	PatientDiagnosisAnamnesisForDelete []string                     `bun:"-" json:"patientDiagnosisAnamnesisForDelete"`
}

func (item *PatientDiagnosis) SetIdForChildren() {
	if len(item.PatientDiagnosisAnamnesis) > 0 {
		for i := range item.PatientDiagnosisAnamnesis {
			item.PatientDiagnosisAnamnesis[i].PatientDiagnosisID = item.ID
		}
	}
}

func GetPatientDiagnosisAnamnesis(items []*PatientDiagnosis) []*PatientDiagnosisAnamnesis {
	itemsForGet := make([]*PatientDiagnosisAnamnesis, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		items[i].SetIdForChildren()
		itemsForGet = append(itemsForGet, items[i].PatientDiagnosisAnamnesis...)
	}
	return itemsForGet
}

func GetPatientDiagnosisAnamnesisForDelete(items []*PatientDiagnosis) []string {
	itemsForGet := make([]string, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		items[i].SetIdForChildren()
		itemsForGet = append(itemsForGet, items[i].PatientDiagnosisAnamnesisForDelete...)
	}
	return itemsForGet
}
