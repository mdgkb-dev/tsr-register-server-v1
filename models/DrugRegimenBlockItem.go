package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugRegimenBlockItem struct {
	bun.BaseModel      `bun:"drug_regimen_block_items,alias:drug_regimen_block_items"`
	ID                 uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	DaysCount          int           `bun:"type:integer" json:"daysCount"`
	OrderItem          int           `bun:"type:integer" json:"orderItem"`
	DrugRegimenBlockID uuid.UUID `bun:"type:uuid" json:"drugRegimenBlockId"`

	DrugRegimenBlock *DrugRegimenBlock `bun:"rel:has-one" json:"drugRegimenBlock"`
}
