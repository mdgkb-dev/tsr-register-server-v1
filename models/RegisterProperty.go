package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterProperty struct {
	bun.BaseModel `bun:"register_property,alias:register_property"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
	ShortName     string    `json:"shortName"`
	ColWidth      string    `json:"colWidth"`
	WithOther     bool      `json:"withOther"`

	ValueType   *ValueType `bun:"rel:belongs-to" json:"valueType"`
	ValueTypeId uuid.UUID  `bun:"type:uuid" json:"valueTypeId"`

	RegisterPropertyRadio          []*RegisterPropertyRadio `bun:"rel:has-many" json:"registerPropertyRadio"`
	RegisterPropertyRadioForDelete []string                 `bun:"-" json:"registerPropertyRadioForDelete"`

	RegisterPropertySet          []*RegisterPropertySet `bun:"rel:has-many" json:"registerPropertySet"`
	RegisterPropertySetForDelete []string               `bun:"-" json:"registerPropertySetForDelete"`
}
