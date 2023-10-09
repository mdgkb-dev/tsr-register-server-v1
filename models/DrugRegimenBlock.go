package models

import (
	"github.com/Pramod-Devireddy/go-exprtk"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugRegimenBlock struct {
	bun.BaseModel `bun:"drug_regimen_blocks,alias:drug_regimen_blocks"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Infinitely    bool      `json:"infinitely"`
	EveryDay      bool      `json:"everyDay"`
	OrderItem     int       `bun:"type:integer" json:"orderItem"`
	DrugRegimenID uuid.UUID `bun:"type:uuid" json:"drugRegimenId"`

	DrugRegimen                    *DrugRegimen          `bun:"rel:belongs-to" json:"drugRegimen"`
	DrugRegimenBlockItems          DrugRegimenBlockItems `bun:"rel:has-many" json:"drugRegimenBlockItems"`
	DrugRegimenBlockItemsForDelete []string              `bun:"-" json:"drugRegimenBlockItemsForDelete"`

	Formula   *Formula      `bun:"rel:belongs-to" json:"formula"`
	FormulaID uuid.NullUUID `bun:"type:uuid" json:"formulaId"`
}

type DrugRegimenBlocks []*DrugRegimenBlock

func (item *DrugRegimenBlock) SetIDForChildren() {
	if len(item.DrugRegimenBlockItems) > 0 {
		for i := range item.DrugRegimenBlockItems {
			item.DrugRegimenBlockItems[i].DrugRegimenBlockID = item.ID
		}
	}
}

func (items DrugRegimenBlocks) CalculateNeeding(variables map[string]interface{}, m exprtk.GoExprtk, daysQuantity uint) float64 {
	sum := float64(0)
	for i := range items {
		sum += items[i].CalculateNeeding(variables, m, daysQuantity)
	}
	return sum
}

func (item *DrugRegimenBlock) CalculateNeeding(variables map[string]interface{}, m exprtk.GoExprtk, daysQuantity uint) float64 {
	blockQuantity := item.Formula.Calculate(variables, m)
	if daysQuantity == 0 {
		return blockQuantity * float64(daysQuantity)
	}
	return float64(daysQuantity) * blockQuantity
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
