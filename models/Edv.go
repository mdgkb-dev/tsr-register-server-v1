package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Edv struct {
	bun.BaseModel `bun:"edv,alias:edv"`
	ID            uuid.UUID   `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Disability    *Disability `bun:"rel:belongs-to" json:"disability"`
	DisabilityID  uuid.UUID   `bun:"type:uuid" json:"disabilityId"`
	Parameter1    bool        `json:"parameter1"`
	Parameter2    bool        `json:"parameter2"`
	Parameter3    bool        `json:"parameter3"`
	Period        *Period     `bun:"rel:belongs-to" json:"period"`
	PeriodID      uuid.UUID   `bun:"type:uuid" json:"periodId"`
}

func GetPeriodsFromEdv(items []*Edv) []*Period {
	itemsForGet := make([]*Period, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].Period)
	}
	return itemsForGet
}

func SetPeriodIDToEdv(items []*Edv) {
	if len(items) == 0 {
		return
	}
	for i := range items {
		items[i].PeriodID = items[i].Period.ID
	}
	return
}
