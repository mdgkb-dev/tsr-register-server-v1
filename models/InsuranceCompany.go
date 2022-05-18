package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type InsuranceCompany struct {
	bun.BaseModel `bun:"insurance_companies,alias:insurance_companies"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
}
