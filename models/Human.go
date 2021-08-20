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

	Documents          []*Document `bun:"rel:has-many" json:"documents"`
	DocumentsForDelete []string    `bun:"-" json:"documentsForDelete"`
}

type InsuranceCompanyToHuman struct {
	bun.BaseModel      `bun:"insurance_company_to_human,alias:insurance_company_to_human"`
	ID                 uuid.UUID         `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Number             int               `json:"number"`
	InsuranceCompany   *InsuranceCompany `bun:"rel:belongs-to" json:"insuranceCompany"`
	InsuranceCompanyID uuid.UUID         `bun:"type:uuid" json:"insuranceCompanyId"`
	Human              *Human            `bun:"rel:has-one" json:"human"`
	HumanID            uuid.UUID         `bun:"type:uuid" json:"humanId"`
}
