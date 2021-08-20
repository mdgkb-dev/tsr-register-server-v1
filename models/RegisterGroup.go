package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterGroup struct {
	bun.BaseModel                            `bun:"register_group,alias:register_group"`
	ID                                       uuid.UUID                          `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                                     string                             `json:"name"`
	RegisterPropertyToRegisterGroup          []*RegisterPropertyToRegisterGroup `bun:"rel:has-many" json:"registerPropertyToRegisterGroup"`
	RegisterPropertyToRegisterGroupForDelete []string                           `bun:"-" json:"registerPropertyToRegisterGroupForDelete"`
}

func (item *RegisterGroup) SetIdForChildren() {
	if len(item.RegisterPropertyToRegisterGroup) == 0 {
		return
	}
	for i := range item.RegisterPropertyToRegisterGroup {
		item.RegisterPropertyToRegisterGroup[i].RegisterGroupID = item.ID
	}
}

type RegisterPropertyToRegisterGroup struct {
	bun.BaseModel      `bun:"register_property_to_register_group,alias:register_property_to_register_group"`
	ID                 uuid.UUID         `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	RegisterGroupID    uuid.UUID         `bun:"type:uuid" json:"registerGroupId"`
	RegisterGroup      *RegisterGroup    `bun:"rel:belongs-to" json:"registerGroup"`
	RegisterPropertyID uuid.UUID         `bun:"type:uuid" json:"registerPropertyId"`
	RegisterProperty   *RegisterProperty `bun:"rel:belongs-to" json:"registerProperty"`
	Order              int               `bun:"order" json:"order"`
}
