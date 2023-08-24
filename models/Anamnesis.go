package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Anamnesis struct {
	bun.BaseModel `bun:"anamneses,alias:anamneses"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	MkbItem       *MkbItem      `bun:"rel:belongs-to" json:"mkbItem"`
	MkbItemID     uuid.NullUUID `bun:"type:uuid" json:"mkbItemId"`
	Value         string        `json:"value"`
	Date          time.Time     `bun:"item_date" json:"date"`
	DoctorName    string        `json:"doctorName"`

	Patient   *Patient      `bun:"rel:belongs-to" json:"patient"`
	PatientID uuid.NullUUID `bun:"type:uuid" json:"patientId"`
}

type Anamneses []*Anamnesis

type AnamnesesWithCount struct {
	Anamneses Anamneses `json:"items"`
	Count     int       `json:"count"`
}
