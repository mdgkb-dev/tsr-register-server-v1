package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugDoze struct {
	bun.BaseModel `bun:"drug_dozes,alias:drug_dozes"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Quantity      float32       `json:"quantity"`

	DrugDozeComponents DrugDozeComponents `bun:"rel:has-many" json:"drugComponents"`
	DrugRegimens       DrugRegimens       `bun:"rel:has-many" json:"drugRegimenBlocks"`

	DrugForm   *DrugForm     `bun:"rel:belongs-to" json:"drugForm"`
	DrugFormID uuid.NullUUID `bun:"type:uuid" json:"drugFormId"`
}
type DrugDozes []*DrugDoze

type DrugDozesWithCount struct {
	DrugDozes DrugDozes `json:"items"`
	Count     int       `json:"count"`
}

func (item *DrugDoze) GetActiveComponentsSum() float64 {
	sum := float64(0)
	for _, component := range item.DrugDozeComponents {
		sum += float64(component.Quantity)
	}
	return sum
}

func (item *DrugDoze) Count() float64 {
	sum := float64(0)
	for _, component := range item.DrugDozeComponents {
		sum += float64(component.Quantity)
	}
	return sum
}
