package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertyMeasure struct {
	bun.BaseModel      `bun:"register_property_measures,alias:register_property_measures"`
	ID                 uuid.UUID         `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name               string            `json:"name"`
	Order              int               `bun:"register_property_measure_order"" json:"order"`
	RegisterPropertyID uuid.UUID         `bun:"type:uuid" json:"registerPropertyId"`
	RegisterProperty   *RegisterProperty `bun:"rel:belongs-to" json:"registerProperty"`
}

type RegisterPropertyMeasures []*RegisterPropertyMeasure
