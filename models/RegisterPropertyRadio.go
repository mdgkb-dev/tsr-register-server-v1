package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterPropertyRadio struct {
	bun.BaseModel                   `bun:"register_property_radio,alias:register_property_radio"`
	ID                              uuid.NullUUID          `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                            string                 `json:"name"`
	RegisterPropertyID              uuid.UUID              `bun:"type:uuid" json:"registerPropertyID"`
	Order                           int                    `bun:"register_property_radio_order" json:"order"`
	RegisterPropertyOthers          RegisterPropertyOthers `bun:"rel:has-many" json:"registerPropertyOthers"`
	RegisterPropertyOthersForDelete []uuid.UUID            `bun:"-" json:"registerPropertyOthersForDelete"`
}

type RegisterPropertyRadios []*RegisterPropertyRadio

func (item *RegisterPropertyRadio) SetIDForChildren() {
	if len(item.RegisterPropertyOthers) == 0 {
		return
	}
	for i := range item.RegisterPropertyOthers {
		item.RegisterPropertyOthers[i].RegisterPropertyRadioID = item.ID
	}
}

func (items RegisterPropertyRadios) SetIDForChildren() {
	if len(items) == 0 {
		return
	}
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (items RegisterPropertyRadios) GetRegisterPropertyOthers() RegisterPropertyOthers {
	itemsForGet := make(RegisterPropertyOthers, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyOthers...)
	}
	return itemsForGet
}

func (items RegisterPropertyRadios) GetRegisterPropertyOthersForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyOthersForDelete...)
	}
	return itemsForGet
}
