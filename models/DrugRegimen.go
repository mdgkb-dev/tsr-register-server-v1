package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugRegimen struct {
	bun.BaseModel `bun:"drug_regimens,alias:drug_regimens"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
	DrugID        uuid.UUID `bun:"type:uuid" json:"drugId"`

	Drug                       *Drug               `bun:"rel:belongs-to" json:"drug"`
	DrugRegimenBlocks          []*DrugRegimenBlock `bun:"rel:has-many" json:"drugRegimenBlocks"`
	DrugRegimenBlocksForDelete []string            `bun:"-" json:"drugRegimenBlocksForDelete"`
}

func (item *DrugRegimen) SetIdForChildren() {
	if len(item.DrugRegimenBlocks) > 0 {
		for i := range item.DrugRegimenBlocks {
			item.DrugRegimenBlocks[i].DrugRegimenID = item.ID
		}
	}
}

func GetDrugRegimenBlocks(items []*DrugRegimen) []*DrugRegimenBlock {
	itemsForGet := make([]*DrugRegimenBlock, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		items[i].SetIdForChildren()
		itemsForGet = append(itemsForGet, items[i].DrugRegimenBlocks...)
	}
	return itemsForGet
}

func GetDrugRegimenBlocksForDelete(items []*DrugRegimen) []string {
	itemsForGet := make([]string, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		items[i].SetIdForChildren()
		itemsForGet = append(itemsForGet, items[i].DrugRegimenBlocksForDelete...)
	}
	return itemsForGet
}
