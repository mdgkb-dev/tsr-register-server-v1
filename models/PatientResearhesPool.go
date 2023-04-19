package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientResearchesPool struct {
	bun.BaseModel `bun:"patients_researches_pools,alias:patients_researches_pools"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Patient       *Patient      `bun:"rel:belongs-to" json:"patient"`
	PatientID     uuid.NullUUID `bun:"type:uuid" json:"patientId"`

	ResearchesPool   *ResearchesPool `bun:"rel:belongs-to" json:"researchesPool"`
	ResearchesPoolID uuid.NullUUID   `bun:"type:uuid" json:"researchesPoolId"`
}

type PatientsResearchesPools []*PatientResearchesPool

type PatientsResearchesPoolsWithCount struct {
	PatientsResearchesPools PatientsResearchesPools `json:"items"`
	Count                   int                     `json:"count"`
}
