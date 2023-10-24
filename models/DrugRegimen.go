package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugRegimen struct {
	bun.BaseModel `bun:"drug_regimens,alias:drug_regimens"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`

	DrugDoze   *DrugDoze     `bun:"rel:belongs-to" json:"drugDoze"`
	DrugDozeID uuid.NullUUID `bun:"type:uuid" json:"drugDozeId"`

	DrugRegimenBlocks          DrugRegimenBlocks `bun:"rel:has-many" json:"drugRegimenBlocks"`
	DrugRegimenBlocksForDelete []string          `bun:"-" json:"drugRegimenBlocksForDelete"`

	MaxMonths *uint `json:"maxMonths"`
	MaxWeight *uint `json:"maxWeight"`
}

type DrugRegimens []*DrugRegimen
type DrugRegimensWithCount struct {
	DrugRegimens DrugRegimens `json:"items"`
	Count        int          `json:"count"`
}

func (item *DrugRegimen) SetIDForChildren() {
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
		items[i].SetIDForChildren()
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
		items[i].SetIDForChildren()
		itemsForGet = append(itemsForGet, items[i].DrugRegimenBlocksForDelete...)
	}
	return itemsForGet
}

func (item *DrugRegimen) CalculateNeeding(variables map[string]interface{}, periodInDays uint, measuresInPack float64) float64 {
	needing := item.DrugRegimenBlocks.CalculateNeeding(variables, periodInDays, measuresInPack)
	fmt.Println("needing:", needing)
	return needing
}

func (items DrugRegimens) FindDrugRegimen(weight uint, months uint) *DrugRegimen {
	for i := range items {
		if items[i].CheckConditions(weight, months) {
			return items[i]
		}
	}
	return nil
}

func (item *DrugRegimen) CheckConditions(weight uint, months uint) bool {
	if item != nil {
		fmt.Println(weight, months, *item.MaxMonths, *item.MaxWeight, item.Name)
	}

	if item.MaxMonths != nil && *item.MaxMonths < months {
		return false
	}

	if item.MaxWeight != nil && *item.MaxWeight < weight {
		return false
	}
	return true
}
