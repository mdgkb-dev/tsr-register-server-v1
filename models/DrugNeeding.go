package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugNeeding struct {
	bun.BaseModel `bun:"drug_needings,alias:drug_needings"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Measures      uint          `json:"measures"`
	Packs         uint          `json:"packs"`
	Calculation   string        `json:"calculation"`
	Weight        uint          `json:"weight"`
	AgeInMonths   uint          `json:"ageInMonths"`

	DrugRegimen   *DrugRegimen  `bun:"rel:belongs-to" json:"drugRegimen"`
	DrugRegimenID uuid.NullUUID `bun:"type:uuid" json:"drugRegimenId"`
}

type DrugNeedings []*DrugNeeding
type DrugNeedingsWithCount struct {
	DrugNeedings DrugNeedings `json:"items"`
	Count        int          `json:"count"`
}

func (item *DrugNeeding) Init(drugRegimen *DrugRegimen, measures uint, packs uint) {
	item.DrugRegimen = drugRegimen
	item.Measures = measures
	item.Packs = packs
	item.Calculation = fmt.Sprintf("%dгр, %dуп.", measures, packs)
}
