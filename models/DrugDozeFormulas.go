package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugRegimenBlockFormula struct {
	bun.BaseModel `bun:"drug_regimen_block_formulas,alias:drug_regimen_block_formulas"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	DrugDoze   *DrugDoze     `bun:"rel:belongs-to" json:"drugDoze"`
	DrugDozeID uuid.NullUUID `bun:"type:uuid" json:"drugDozeId"`
}

type DrugRegimenBlockFormulas []*DrugRegimenBlockFormula

type DrugDozeFormulasWithCount struct {
	DrugComponents DrugDozeComponents `json:"items"`
	Count          int                `json:"count"`
}
