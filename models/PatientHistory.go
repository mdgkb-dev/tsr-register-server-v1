package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientHistory struct {
	bun.BaseModel    `bun:"patient_histories,alias:patient_histories"`
	PatientHistoryID uuid.UUID `bun:"type:uuid,nullzero,notnull,default:uuid_generate_v4()" json:"patientHistoryId" `
	History          *History  `bun:"rel:belongs-to" json:"history"`
	HistoryID        uuid.UUID `bun:"type:uuid" json:"historyId"`
	Patient
}
