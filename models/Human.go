package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Human struct {
	bun.BaseModel       `bun:"human,alias:human"`
	ID                  uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                string    `json:"name"`
	Surname             string    `json:"surname"`
	Patronymic          string    `json:"patronymic"`
	IsMale              bool      `json:"isMale"`
	DateBirth           time.Time `json:"dateBirth"`
	AddressRegistration string    `json:"addressRegistration"`
	AddressResidential  string    `json:"addressResidential"`
	Contact             *Contact  `bun:"rel:belongs-to" json:"contact"`
	ContactID           uuid.UUID `bun:"type:uuid" json:"contactId"`

	Documents          []*Document `bun:"rel:has-many" json:"documents"`
	DocumentsForDelete []string    `bun:"-" json:"documentsForDelete"`

	InsuranceCompanyToHuman          []*InsuranceCompanyToHuman `bun:"rel:has-many" json:"insuranceCompanyToHuman"`
	InsuranceCompanyToHumanForDelete []string                   `bun:"-" json:"insuranceCompanyToHumanForDelete"`
}

func (item *Human) SetIdForChildren() {
	if len(item.Documents) > 0 {
		for i := range item.Documents {
			item.Documents[i].HumanID = item.ID
		}
	}
	if len(item.InsuranceCompanyToHuman) > 0 {
		for i := range item.InsuranceCompanyToHuman {
			item.InsuranceCompanyToHuman[i].HumanID = item.ID
		}
	}
}

type InsuranceCompanyToHuman struct {
	bun.BaseModel      `bun:"insurance_company_to_human,alias:insurance_company_to_human"`
	ID                 uuid.UUID         `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Number             string            `json:"number"`
	InsuranceCompany   *InsuranceCompany `bun:"rel:belongs-to" json:"insuranceCompany"`
	InsuranceCompanyID uuid.UUID         `bun:"type:uuid" json:"insuranceCompanyId"`
	Human              *Human            `bun:"rel:belongs-to" json:"human"`
	HumanID            uuid.UUID         `bun:"type:uuid" json:"humanId"`
}
