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
