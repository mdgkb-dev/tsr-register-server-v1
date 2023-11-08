package models

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/middleware"

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

func CreateFromPatientID(c context.Context, patientID uuid.NullUUID) (PatientsDomains, error) {
	domainsIDS := middleware.ClaimDomainIDS.FromContextSlice(c)
	items := make(PatientsDomains, len(domainsIDS)-1)
	for i := range domainsIDS {
		dID, err := uuid.Parse(domainsIDS[i])
		if err != nil {
			return nil, err
		}
		d := PatientDomain{PatientID: patientID, DomainID: uuid.NullUUID{UUID: dID, Valid: true}}
		items[i] = &d
	}
	return items, nil
}
