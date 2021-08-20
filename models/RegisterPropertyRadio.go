package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertyRadio struct {
	bun.BaseModel      `bun:"register_property_radio,alias:register_property_radio"`
	ID                 uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name               string    `json:"name"`
	RegisterPropertyID uuid.UUID `bun:"type:uuid" json:"registerPropertyID"`
}
