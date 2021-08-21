package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Disability struct {
	bun.BaseModel `bun:"disability,alias:disability"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Patient       *Patient  `bun:"rel:belongs-to" json:"patient"`
	PatientID     uuid.UUID `bun:"type:uuid" json:"patientId"`
	Edvs          []*Edv    `bun:"rel:has-many" json:"edvs"`
	EdvsForDelete []string  `bun:"-" json:"edvsForDelete"`
	Period        *Period   `bun:"rel:belongs-to" json:"period"`
	PeriodID      uuid.UUID `bun:"type:uuid" json:"periodId"`
}

func (item *Disability) SetIdForChildren() {
	if len(item.Edvs) > 0 {
		for i := range item.Edvs {
			item.Edvs[i].DisabilityID = item.ID
		}
	}
}

func GetEdvs(disabilities []*Disability) []*Edv {
	items := make([]*Edv, 0)
	if len(disabilities) == 0 {
		return items
	}
	for i := range disabilities {
		disabilities[i].SetIdForChildren()
		items = append(items, disabilities[i].Edvs...)
	}
	return items
}

func GetPeriodsFromDisability(items []*Disability) []*Period {
	itemsForGet := make([]*Period, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].Period)
	}
	return itemsForGet
}

func SetPeriodIDToDisabilities(items []*Disability) {
	if len(items) == 0 {
		return
	}
	for i := range items {
		items[i].PeriodID = items[i].Period.ID
	}
	return
}
