package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type HeightWeight struct {
	bun.BaseModel `bun:"height_weight,alias:height_weight"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Height        int       `json:"height"`
	Weight        int       `json:"weight"`
	Date          time.Time `json:"date"`
	Patient       *Patient  `bun:"rel:belongs-to" json:"patients"`
	PatientID     uuid.UUID `bun:"type:uuid" json:"patientId"`
	DeletedAt     time.Time `bun:",soft_delete" json:"deletedAt"`
}
