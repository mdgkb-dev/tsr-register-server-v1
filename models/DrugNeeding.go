package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugNeeding struct {
	bun.BaseModel `bun:"drugs,alias:drugs"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	DrugRegimen   *DrugRegimen  `bun:"rel:has-many" json:"drugRegimen"`
	Measures      uint          `json:"measures"`
	Packs         uint          `json:"packs"`
	Calculation   string        `json:"calculation"`
}

type DrugNeedings []*DrugNeeding
type DrugNeedingsWithCount struct {
	Drugs Drugs `json:"items"`
	Count int   `json:"count"`
}

func (item *DrugNeeding) Init(drugRegimen *DrugRegimen, measures uint, packs uint, calculation string) {
	item.DrugRegimen = drugRegimen
	item.Measures = measures
	item.Packs = packs
	item.Calculation = calculation
}
