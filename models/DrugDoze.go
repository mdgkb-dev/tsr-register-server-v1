package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugDoze struct {
	bun.BaseModel `bun:"drug_dozes,alias:drugs"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
	Quantity      float32   `json:"quantity"`

	DrugForm   *DrugForm     `bun:"rel:belongs-to" json:"drugForm"`
	DrugFormID uuid.NullUUID `bun:"type:uuid" json:"drugFormId"`
}
type DrugDozes []*DrugDoze

type DrugDozesWithCount struct {
	DrugDozes DrugDozes `json:"items"`
	Count     int       `json:"count"`
}
