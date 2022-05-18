package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertyOther struct {
	bun.BaseModel `bun:"register_property_others,alias:register_property_others"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
	//ShortName     string    `json:"shortName"`
	//ColWidth      string    `json:"colWidth"`

	RegisterProperty   *RegisterProperty `bun:"rel:belongs-to" json:"registerProperty"`
	RegisterPropertyID uuid.NullUUID     `bun:"type:uuid" json:"registerPropertyId"`

	RegisterPropertyRadio   *RegisterPropertyRadio `bun:"rel:belongs-to" json:"registerPropertyRadio"`
	RegisterPropertyRadioID uuid.NullUUID          `bun:"type:uuid" json:"registerPropertyRadioId"`

	RegisterPropertySet   *RegisterPropertyRadio `bun:"rel:belongs-to" json:"registerPropertySet"`
	RegisterPropertySetID uuid.NullUUID          `bun:"type:uuid" json:"registerPropertySetId"`
	Order                 int                    `bun:"register_property_others_order" json:"order"`
	//RegisterPropertyRadio          []*RegisterPropertyRadio `bun:"rel:has-many" json:"registerPropertyRadio"`
	//RegisterPropertyRadioForDelete []string                 `bun:"-" json:"registerPropertyRadioForDelete"`
	//
	//RegisterPropertySet          []*RegisterPropertySet `bun:"rel:has-many" json:"registerPropertySet"`
	//RegisterPropertySetForDelete []string               `bun:"-" json:"registerPropertySetForDelete"`
	//
	//RegisterPropertyToRegisterGroup []*RegisterPropertyToRegisterGroup `bun:"rel:has-many" json:"registerPropertyToRegisterGroup"`
}

type RegisterPropertyOthers []*RegisterPropertyOther
