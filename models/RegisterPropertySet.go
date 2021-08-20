package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertySet struct {
	bun.BaseModel      `bun:"register_property_set,alias:register_property_set"`
	ID                 uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name               string    `json:"name"`
	RegisterPropertyID uuid.UUID `bun:"type:uuid" json:"registerPropertyID"`
}
