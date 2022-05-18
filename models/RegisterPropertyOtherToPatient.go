package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertyOtherToPatient struct {
	bun.BaseModel           `bun:"register_property_other_to_patient,alias:register_property_other_to_patient"`
	ID                      uuid.UUID              `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Value                   string                 `json:"value"`
	RegisterPropertyOther   *RegisterPropertyOther `bun:"rel:belongs-to" json:"registerProperty"`
	RegisterPropertyOtherID uuid.UUID              `bun:"type:uuid" json:"registerPropertyOtherId"`
	Patient                 *Patient               `bun:"rel:has-one" json:"patients"`
	PatientID               uuid.UUID              `bun:"type:uuid" json:"patientId"`
}

type RegisterPropertyOthersToPatient []*RegisterPropertyOtherToPatient
