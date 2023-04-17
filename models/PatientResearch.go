package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientResearch struct {
	bun.BaseModel `bun:"research_results,alias:research_results"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Research        *Research       `bun:"rel:belongs-to" json:"research"`
	ResearchID      uuid.UUID       `bun:"type:uuid" json:"researchId"`
	Patient         *Patient        `bun:"rel:belongs-to" json:"patients"`
	PatientID       uuid.UUID       `bun:"type:uuid" json:"patientId"`
	ResearchResults ResearchResults `bun:"rel:has-many" json:"researchResults"`

	Order uint `bun:"item_order" json:"order"`
}

type PatientResearches []*PatientResearch
