package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FormulaResult struct {
	bun.BaseModel `bun:"formula_results,alias:formula_results"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string        `json:"name"`
	LowRange      float64       `json:"lowRange"`
	HighRange     float64       `json:"highRange"`

	Formula   *Formula      `bun:"rel:belongs-to" json:"formula"`
	FormulaID uuid.NullUUID `bun:"type:uuid" json:"formulaId"`
}

type FormulaResults []*FormulaResult
