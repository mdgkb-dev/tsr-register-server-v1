package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type AnthropometryData struct {
	bun.BaseModel   `bun:"anthropometry_data,alias:anthropometry_data"`
	ID              uuid.UUID      `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Anthropometry   *Anthropometry `bun:"rel:belongs-to" json:"anthropometry"`
	AnthropometryID uuid.UUID      `bun:"type:uuid" json:"anthropometryId"`
	Patient         *Patient       `bun:"rel:belongs-to" json:"patient"`
	PatientID       uuid.UUID      `bun:"type:uuid" json:"patientId"`
	Value           int            `json:"value"`
	Date            time.Time      `json:"date"`
}
