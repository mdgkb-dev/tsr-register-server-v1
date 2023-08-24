package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type AnamnesisResearchDomain struct {
	bun.BaseModel `bun:"anamnesis_researches_domains,alias:anamnesis_researches_domains"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	ResearchID    uuid.NullUUID `bun:"type:uuid" json:"researchId"`
	Research      *Research     `bun:"-" json:"research"`

	DomainID uuid.NullUUID `bun:"type:uuid" json:"domainId"`
	Domain   *Domain       `bun:"-" json:"domain"`
}

type AnamnesisResearchesDomains []*Research
