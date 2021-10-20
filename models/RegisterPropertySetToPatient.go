package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertySetToPatient struct {
	bun.BaseModel         `bun:"register_property_set_to_patient,alias:register_property_set_to_patient"`
	ID                    uuid.UUID            `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	RegisterPropertySet   *RegisterPropertySet `bun:"rel:belongs-to" json:"registerPropertySet"`
	RegisterPropertySetID uuid.UUID            `bun:"type:uuid" json:"registerPropertySetId"`
	Patient               *Patient             `bun:"rel:has-one" json:"patient"`
	PatientID             uuid.UUID            `bun:"type:uuid" json:"patientId"`
	DeletedAt             time.Time            `bun:",soft_delete" json:"deletedAt"`
}