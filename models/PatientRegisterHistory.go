package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientRegisterHistory struct {
	bun.BaseModel     `bun:"patients_registers,alias:patients_registers"`
	ID                uuid.UUID        `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	PatientRegister   *PatientRegister `bun:"rel:has-one" json:"patientRegister"`
	PatientRegisterID uuid.UUID        `bun:"type:uuid" json:"patientRegisterId"`
	Date              *time.Time       `bun:"item_date" json:"date"`
	User              *User            `bun:"rel:has-one" json:"user"`
	UserID            uuid.UUID        `bun:"type:uuid" json:"userId"`
}

type PatientRegisterHistories []*PatientRegisterHistory
