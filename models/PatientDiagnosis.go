package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientDiagnosis struct {
	bun.BaseModel `bun:"patient_diagnosis,alias:patient_diagnosis"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Patient       *Patient      `bun:"rel:has-one" json:"patients"`
	PatientID     uuid.NullUUID `bun:"type:uuid" json:"PatientId"`
	DoctorName    string        `json:"doctorName"`

	Primary                            bool          `json:"primary"`
	DeletedAt                          *time.Time    `bun:",soft_delete" json:"deletedAt"`
	MkbItem                            *MkbItem      `bun:"rel:belongs-to" json:"mkbItem"`
	MkbItemID                          uuid.NullUUID `bun:"type:uuid" json:"mkbItemId"`
	Anamneses                          Anamneses     `bun:"rel:has-many" json:"anamneses"`
	PatientDiagnosisAnamnesisForDelete []string      `bun:"-" json:"patientDiagnosisAnamnesisForDelete"`
}

type PatientDiagnosisWithCount struct {
	PatientDiagnosis []*PatientDiagnosis `json:"items"`
	Count            int                 `json:"count"`
}

func (item *PatientDiagnosis) SetIDForChildren() {
	if len(item.Anamneses) > 0 {
		for i := range item.Anamneses {
			item.Anamneses[i].PatientDiagnosisID = item.ID
		}
	}
}

func GetPatientDiagnosisAnamnesis(items []*PatientDiagnosis) Anamneses {
	itemsForGet := make(Anamneses, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		items[i].SetIDForChildren()
		itemsForGet = append(itemsForGet, items[i].Anamneses...)
	}
	return itemsForGet
}

func GetPatientDiagnosisAnamnesisForDelete(items []*PatientDiagnosis) []string {
	itemsForGet := make([]string, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		items[i].SetIDForChildren()
		itemsForGet = append(itemsForGet, items[i].PatientDiagnosisAnamnesisForDelete...)
	}
	return itemsForGet
}
