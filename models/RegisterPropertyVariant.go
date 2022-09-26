package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertyVariant struct {
	bun.BaseModel      `bun:"register_property_variants,alias:register_property_variants"`
	ID                 uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name               string        `json:"name"`
	RegisterPropertyID uuid.UUID     `bun:"type:uuid" json:"registerPropertyID"`
}

type RegisterPropertyVariants []*RegisterPropertyVariant
