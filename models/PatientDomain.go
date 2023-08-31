package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PatientDomain struct {
	bun.BaseModel `bun:"patients_domains,alias:patients_domains"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	PatientID     uuid.NullUUID `bun:"type:uuid" json:"patientId"`
	Patient       *Patient      `bun:"-" json:"patient"`

	DomainID uuid.NullUUID `bun:"type:uuid" json:"domainId"`
	Domain   *Domain       `bun:"-" json:"domain"`
}

type PatientsDomains []*PatientDomain

type PatientsDomainsWithCount struct {
	PatientsDomains PatientsDomains `json:"items"`
	Count           int             `json:"count"`
}
