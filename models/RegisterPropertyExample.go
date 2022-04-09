package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertyExample struct {
	bun.BaseModel      `bun:"register_property_examples,alias:register_property_examples"`
	ID                 uuid.UUID         `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name               string            `json:"name"`
	RegisterPropertyID uuid.UUID         `bun:"type:uuid" json:"registerPropertyId"`
	RegisterProperty   *RegisterProperty `bun:"rel:belongs-to" json:"registerProperty"`
}

type RegisterPropertyExamples []*RegisterPropertyExample
