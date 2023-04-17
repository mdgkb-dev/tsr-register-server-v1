package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientRegister struct {
	bun.BaseModel `bun:"patients_registers,alias:patients_registers"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Patient       *Patient  `bun:"rel:has-one" json:"patients"`
	PatientID     uuid.UUID `bun:"type:uuid" json:"PatientId"`

	Register       *Patient  `bun:"rel:has-one" json:"register"`
	RegisterID     uuid.UUID `bun:"type:uuid" json:"registerID"`
}

type PatientsRegisters []*PatientRegister