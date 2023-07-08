package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugForm struct {
	bun.BaseModel    `bun:"drug_forms,alias:drugs"`
	ID               uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name             string        `json:"name"`
	ReportName       string        `json:"reportName"`
	DateRegistration *time.Time    `json:"dateRegistration"`
	DrugDozes        DrugDozes     `bun:"rel:has-many" json:"drugDozes"`

	Drug   *Drug         `bun:"rel:belongs-to" json:"drug"`
	DrugID uuid.NullUUID `bun:"type:uuid" json:"drugId"`
}

type DrugForms []*DrugForm

type DrugFormsWithCount struct {
	DrugForms DrugForms `json:"items"`
	Count     int       `json:"count"`
}
