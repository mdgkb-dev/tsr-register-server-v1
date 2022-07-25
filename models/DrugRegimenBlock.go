package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugRegimenBlock struct {
	bun.BaseModel `bun:"drug_regimen_blocks,alias:drug_regimen_blocks"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Infinitely    bool      `json:"infinitely"`
	OrderItem     int       `bun:"type:integer" json:"orderItem"`
	DrugRegimenID uuid.UUID `bun:"type:uuid" json:"drugRegimenId"`

	DrugRegimen                    *DrugRegimen            `bun:"rel:belongs-to" json:"drugRegimen"`
	DrugRegimenBlockItems          []*DrugRegimenBlockItem `bun:"rel:has-many" json:"drugRegimenBlockItems"`
	DrugRegimenBlockItemsForDelete []string                `bun:"-" json:"drugRegimenBlockItemsForDelete"`
}

func (item *DrugRegimenBlock) SetIDForChildren() {
	if len(item.DrugRegimenBlockItems) > 0 {
		for i := range item.DrugRegimenBlockItems {
			item.DrugRegimenBlockItems[i].DrugRegimenBlockID = item.ID
		}
	}
}

func GetDrugRegimenBlockItems(items []*DrugRegimenBlock) []*DrugRegimenBlockItem {
	itemsForGet := make([]*DrugRegimenBlockItem, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		items[i].SetIDForChildren()
		itemsForGet = append(itemsForGet, items[i].DrugRegimenBlockItems...)
	}
	return itemsForGet
}

func GetDrugRegimenBlockItemsForDelete(items []*DrugRegimenBlock) []string {
	itemsForGet := make([]string, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		items[i].SetIDForChildren()
		itemsForGet = append(itemsForGet, items[i].DrugRegimenBlockItemsForDelete...)
	}
	return itemsForGet
}
