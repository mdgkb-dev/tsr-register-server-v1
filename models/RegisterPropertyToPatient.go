package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertyToPatient struct {
	bun.BaseModel `bun:"register_property_to_patient,alias:register_property_to_patient"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `

	ValueString string       `json:"valueString"`
	ValueNumber int       `json:"valueNumber"`
	ValueDate   time.Time `json:"valueDate"`
	ValueOther  string    `json:"valueOther"`

	RegisterPropertyRadio   *RegisterPropertyRadio `bun:"rel:belongs-to" json:"registerPropertyRadio"`
	RegisterPropertyRadioID uuid.NullUUID          `bun:"type:uuid" json:"registerPropertyRadioId"`

	RegisterProperty   *RegisterProperty `bun:"rel:belongs-to" json:"registerProperty"`
	RegisterPropertyID uuid.UUID         `bun:"type:uuid" json:"registerPropertyId"`
	Patient            *Patient          `bun:"rel:has-one" json:"patient"`
	PatientID          uuid.UUID         `bun:"type:uuid" json:"patientId"`
	DeletedAt          time.Time         `bun:",soft_delete" json:"deletedAt"`
}
