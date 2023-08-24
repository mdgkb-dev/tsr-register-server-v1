package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type AnamnesisResearch struct {
	bun.BaseModel `bun:"anamneses_researches,alias:anamneses_researches"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	ResearchID    uuid.NullUUID `bun:"type:uuid" json:"researchId"`
	Research      *Research     `bun:"-" json:"research"`

	DomainID uuid.NullUUID `bun:"type:uuid" json:"domainId"`
	Domain   *Domain       `bun:"-" json:"domain"`
}

type AnamnesesResearches []*AnamnesisResearch

type AnamnesesResearchesWithCount struct {
	AnamnesesResearches AnamnesesResearches `json:"items"`
	Count               int                 `json:"count"`
}
