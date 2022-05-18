package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertySet struct {
	bun.BaseModel                   `bun:"register_property_set,alias:register_property_set"`
	ID                              uuid.NullUUID          `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                            string                 `json:"name"`
	RegisterPropertyID              uuid.UUID              `bun:"type:uuid" json:"registerPropertyID"`
	Order                           int                    `bun:"register_property_set_order" json:"order"`
	RegisterPropertyOthers          RegisterPropertyOthers `bun:"rel:has-many" json:"registerPropertyOthers"`
	RegisterPropertyOthersForDelete []uuid.UUID            `bun:"-" json:"registerPropertyOthersForDelete"`
}

type RegisterPropertySets []*RegisterPropertySet

func (item *RegisterPropertySet) SetIdForChildren() {
	if len(item.RegisterPropertyOthers) == 0 {
		return
	}
	for i := range item.RegisterPropertyOthers {
		item.RegisterPropertyOthers[i].RegisterPropertySetID = item.ID
	}
}

func (items RegisterPropertySets) SetIdForChildren() {
	if len(items) == 0 {
		return
	}
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (items RegisterPropertySets) GetRegisterPropertyOthers() RegisterPropertyOthers {
	itemsForGet := make(RegisterPropertyOthers, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyOthers...)
	}
	return itemsForGet
}

func (items RegisterPropertySets) GetRegisterPropertyOthersForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyOthersForDelete...)
	}
	return itemsForGet
}
