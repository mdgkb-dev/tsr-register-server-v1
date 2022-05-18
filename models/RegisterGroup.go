package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterGroup struct {
	bun.BaseModel               `bun:"register_group,alias:register_group"`
	ID                          uuid.UUID          `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                        string             `json:"name"`
	Order                       int                `bun:"register_group_order" json:"order"`
	RegisterProperties          RegisterProperties `bun:"rel:has-many" json:"registerProperties"`
	RegisterPropertiesForDelete []uuid.UUID        `bun:"-" json:"registerPropertiesForDelete"`
	Register                    *Register          `bun:"rel:belongs-to" json:"register"`
	RegisterID                  uuid.UUID          `bun:"type:uuid" json:"registerId"`
}

type RegisterGroups []*RegisterGroup

func (item *RegisterGroup) SetIdForChildren() {
	if len(item.RegisterProperties) == 0 {
		return
	}
	for i := range item.RegisterProperties {
		item.RegisterProperties[i].RegisterGroupID = item.ID
	}
}

func (items RegisterGroups) SetIdForChildren() {
	if len(items) == 0 {
		return
	}
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (items RegisterGroups) GetRegisterProperties() RegisterProperties {
	itemsForGet := make(RegisterProperties, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterProperties...)
	}
	return itemsForGet
}

func (items RegisterGroups) GetRegisterPropertiesForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertiesForDelete...)
	}
	return itemsForGet
}


