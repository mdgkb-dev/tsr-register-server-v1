package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientDrugRegimenItem struct {
	bun.BaseModel        `bun:"patient_drug_regimen_items,alias:patient_drug_regimen_items"`
	ID                   uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Date                 time.Time `json:"date"`
	GettingDate          time.Time `json:"gettingDate"`
	PatientDrugRegimenId uuid.UUID `bun:"type:uuid" json:"patientDrugRegimenId"`

	PatientDrugRegimen *PatientDrugRegimen `bun:"rel:belongs-to" json:"patientDrugRegimen"`
}
