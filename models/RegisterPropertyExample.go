package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertyExample struct {
	bun.BaseModel      `bun:"register_property_examples,alias:register_property_examples"`
	ID                 uuid.UUID         `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name               string            `json:"name"`
	RegisterPropertyID uuid.UUID         `bun:"type:uuid" json:"registerPropertyId"`
	RegisterProperty   *RegisterProperty `bun:"rel:belongs-to" json:"registerProperty"`
	Order              int               `bun:"register_property_example_order" json:"order"`
}

type RegisterPropertyExamples []*RegisterPropertyExample
