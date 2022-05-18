package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientDiagnosisAnamnesis struct {
	bun.BaseModel      `bun:"patient_diagnosis_anamnesis,alias:patient_diagnosis_anamnesis"`
	ID                 uuid.UUID         `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	PatientDiagnosis   *PatientDiagnosis `bun:"rel:belongs-to" json:"patientDiagnosis"`
	PatientDiagnosisID uuid.UUID         `bun:"type:uuid" json:"patientDiagnosisId"`
	Value              string            `json:"value"`
	Date               time.Time         `json:"date"`
}
