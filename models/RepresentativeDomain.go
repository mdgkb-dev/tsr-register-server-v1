package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RepresentativeDomain struct {
	bun.BaseModel `bun:"representatives_domains,alias:representatives_domains"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	RepresentativeID     uuid.NullUUID `bun:"type:uuid" json:"RepresentativeId"`
	Representative       *Representative      `bun:"-" json:"Representative"`

	DomainID uuid.NullUUID `bun:"type:uuid" json:"domainId"`
	Domain   *Domain       `bun:"-" json:"domain"`
}

type RepresentativesDomains []*RepresentativeDomain

type RepresentativesDomainsWithCount struct {
	RepresentativesDomains RepresentativesDomains `json:"items"`
	Count           int             `json:"count"`
}
