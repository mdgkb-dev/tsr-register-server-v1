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
	Tag string `json:"tag"`
	Order         int `bun:"register_property_order" json:"order"`
	ValueType   *ValueType `bun:"rel:belongs-to" json:"valueType"`
	ValueTypeId uuid.UUID  `bun:"type:uuid" json:"valueTypeId"`
	RegisterGroupID    uuid.UUID         `bun:"type:uuid" json:"registerGroupId"`
	RegisterGroup      *RegisterGroup    `bun:"rel:belongs-to" json:"registerGroup"`

	RegisterPropertyRadios          RegisterPropertyRadios `bun:"rel:has-many" json:"registerPropertyRadios"`
	RegisterPropertyRadiosForDelete []uuid.UUID                 `bun:"-" json:"registerPropertyRadiosForDelete"`

	RegisterPropertySets          RegisterPropertySets `bun:"rel:has-many" json:"registerPropertySets"`
	RegisterPropertySetsForDelete []uuid.UUID               `bun:"-" json:"registerPropertySetsForDelete"`
}

type RegisterProperties []*RegisterProperty

func (item *RegisterProperty) SetIdForChildren() {
	if len(item.RegisterPropertyRadios) > 0 {

	for i := range item.RegisterPropertyRadios {
		item.RegisterPropertyRadios[i].RegisterPropertyID = item.ID
	}
	}
	if len(item.RegisterPropertySets) > 0 {
	for i := range item.RegisterPropertySets {
		item.RegisterPropertySets[i].RegisterPropertyID = item.ID
	}
	}
}

func (items RegisterProperties) SetIdForChildren() {
	if len(items) == 0 {
		return
	}
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (items RegisterProperties) GetRegisterPropertyRadios()RegisterPropertyRadios {
	itemsForGet := make(RegisterPropertyRadios, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyRadios...)
	}
	return itemsForGet
}

func (items RegisterProperties) GetRegisterPropertySets()RegisterPropertySets {
	itemsForGet := make(RegisterPropertySets, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertySets...)
	}
	return itemsForGet
}



func (items RegisterProperties) GetRegisterPropertyRadioForDelete()[]uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertyRadiosForDelete...)
	}
	return itemsForGet
}

func (items RegisterProperties) GetRegisterPropertySetForDelete()[]uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].RegisterPropertySetsForDelete...)
	}
	return itemsForGet
}
