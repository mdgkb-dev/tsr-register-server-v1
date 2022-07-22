package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertyToPatient struct {
	bun.BaseModel `bun:"register_property_to_patient,alias:register_property_to_patient"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	ValueString string    `json:"valueString"`
	ValueNumber int       `json:"valueNumber"`
	ValueDate   time.Time `json:"valueDate"`
	ValueOther  string    `json:"valueOther"`

	RegisterPropertyRadio   *RegisterPropertyRadio `bun:"rel:belongs-to" json:"registerPropertyRadio"`
	RegisterPropertyRadioID uuid.NullUUID          `bun:"type:uuid" json:"registerPropertyRadioId"`

	RegisterPropertyMeasure   *RegisterPropertyMeasure `bun:"rel:belongs-to" json:"registerPropertyMeasure"`
	RegisterPropertyMeasureID uuid.NullUUID            `bun:"type:uuid" json:"registerPropertyMeasureId"`

	RegisterProperty   *RegisterProperty `bun:"rel:belongs-to" json:"registerProperty"`
	RegisterPropertyID uuid.UUID         `bun:"type:uuid" json:"registerPropertyId"`

	RegisterGroupToPatient   *RegisterGroupToPatient `bun:"rel:belongs-to" json:"registerGroupToPatient"`
	RegisterGroupToPatientID uuid.UUID               `bun:"type:uuid" json:"registerGroupToPatientID"`
}

type RegisterPropertiesToPatients []*RegisterPropertyToPatient
