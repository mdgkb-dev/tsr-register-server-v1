package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Anamnesis struct {
	bun.BaseModel      `bun:"anamneses,alias:anamneses"`
	ID                 uuid.UUID         `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	PatientDiagnosis   *PatientDiagnosis `bun:"rel:belongs-to" json:"patientDiagnosis"`
	PatientDiagnosisID uuid.UUID         `bun:"type:uuid" json:"patientDiagnosisId"`
	Value              string            `json:"value"`
	Date               time.Time         `bun:"item_date" json:"date"`
	DoctorName         string            `json:"doctorName"`
}

type Anamneses []*Anamnesis

type AnamnesesWithCount struct {
	Anamneses Anamneses `json:"items"`
	Count     int       `json:"count"`
}
