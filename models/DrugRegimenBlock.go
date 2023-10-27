package models

import (
	"fmt"

	"github.com/Pramod-Devireddy/go-exprtk"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugRegimenBlock struct {
	bun.BaseModel `bun:"drug_regimen_blocks,alias:drug_regimen_blocks"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	OrderItem     int       `bun:"type:integer" json:"orderItem"`
	DrugRegimenID uuid.UUID `bun:"type:uuid" json:"drugRegimenId"`

	DrugRegimen                    *DrugRegimen          `bun:"rel:belongs-to" json:"drugRegimen"`
	DrugRegimenBlockItems          DrugRegimenBlockItems `bun:"rel:has-many" json:"drugRegimenBlockItems"`
	DrugRegimenBlockItemsForDelete []string              `bun:"-" json:"drugRegimenBlockItemsForDelete"`

	Formula   *Formula      `bun:"rel:belongs-to" json:"formula"`
	FormulaID uuid.NullUUID `bun:"type:uuid" json:"formulaId"`

	DaysCount   int  `bun:"type:integer" json:"daysCount"`
	TimesPerDay int  `bun:"type:integer" json:"timesPerDay"`
	Infinitely  bool `json:"infinitely"`
	EveryDay    bool `json:"everyDay"`
}

type DrugRegimenBlocks []*DrugRegimenBlock

func (item *DrugRegimenBlock) SetIDForChildren() {
	if len(item.DrugRegimenBlockItems) > 0 {
		for i := range item.DrugRegimenBlockItems {
			item.DrugRegimenBlockItems[i].DrugRegimenBlockID = item.ID
		}
	}
}

func (items DrugRegimenBlocks) CalculateNeeding(variables map[string]interface{}, daysQuantity uint, measuresInPack float64) float64 {
	sum := float64(0)
	fmt.Println("days!!!", daysQuantity, variables)
	m := exprtk.NewExprtk()
	daysLeast := int(daysQuantity)

	for _, block := range items {
		// случай, когда дней больше, чем приёма
		if block.Infinitely {

			for {
				// fmt.Println(measuresInPack);
				dl, s := block.CalculateNeeding(variables, m, daysLeast, measuresInPack)
				daysLeast = dl
				sum += s
				if daysLeast <= 0 {
					break
				}
			}
			break
		}

		daysLeast, s := block.CalculateNeeding(variables, m, daysLeast, measuresInPack)
		sum += s
		fmt.Println(daysLeast)
		if daysLeast <= 0 { // Если дней приёма не осталось
			break
		}
	}
	return sum
}

func (item *DrugRegimenBlock) CalculateNeeding(variables map[string]interface{}, m exprtk.GoExprtk, daysQuantity int, measuresInPack float64) (int, float64) {
	blockQuantity := item.Formula.Calculate(variables, m)
	if daysQuantity == 0 {
		return 0, blockQuantity * float64(daysQuantity)
	}
	daysLeast := 0
	// fmt.Println(daysQuantity, item.DaysCount)
	calculation := float64(0)
	if (item.DaysCount > 1 && !item.EveryDay) || item.Infinitely {
		calculation = measuresInPack / blockQuantity
		fmt.Println("calc1", item.EveryDay, item.Infinitely)
		daysLeast = daysQuantity - int(item.DaysCount)
	}
	if daysQuantity > 1 && item.EveryDay && item.Infinitely {
		calculation = float64(daysQuantity) / (measuresInPack / blockQuantity)
		fmt.Println("calc2", daysQuantity, measuresInPack, blockQuantity, calculation)
		daysLeast = 0
	}

	return daysLeast, calculation
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
