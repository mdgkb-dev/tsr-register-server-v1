package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientRegister struct {
	bun.BaseModel `bun:"patients_registers,alias:patients_registers"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Patient       *Patient      `bun:"rel:belongs-to" json:"patient"`
	PatientID     uuid.NullUUID `bun:"type:uuid" json:"patientId"`

	Register   *Register     `bun:"rel:belongs-to" json:"register"`
	RegisterID uuid.NullUUID `bun:"type:uuid" json:"registerId"`

	User   *User         `bun:"rel:belongs-to" json:"user"`
	UserID uuid.NullUUID `bun:"type:uuid" json:"userId"`
}

type PatientsRegisters []*PatientRegister

type PatientsRegistersWithCount struct {
	PatientsRegisters PatientsRegisters `json:"items"`
	Count             int               `json:"count"`
}
