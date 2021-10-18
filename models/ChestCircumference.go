package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ChestCircumference struct {
	bun.BaseModel `bun:"chest_circumference,alias:chest_circumference"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Value         int           `json:"value"`
	Date          time.Time     `json:"date"`
	DeletedAt     time.Time     `bun:",soft_delete" json:"deletedAt"`
	Patient       *Patient      `bun:"rel:belongs-to" json:"patient"`
	PatientID     uuid.UUID     `bun:"type:uuid" json:"patientId"`
}
