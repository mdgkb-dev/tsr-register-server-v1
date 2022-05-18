package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterToPatient struct {
	bun.BaseModel `bun:"register_to_patient,alias:register_to_patient"`
	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Register      *Register  `bun:"rel:belongs-to" json:"register"`
	RegisterID    uuid.UUID  `bun:"type:uuid" json:"registerId"`
	Patient       *Patient   `bun:"rel:belongs-to" json:"patient"`
	PatientID     uuid.UUID  `bun:"type:uuid" json:"patientId"`
	DeletedAt     *time.Time `bun:",soft_delete" json:"deletedAt"`
}
