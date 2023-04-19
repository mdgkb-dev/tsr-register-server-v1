package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ChopScaleTest struct {
	bun.BaseModel                 `bun:"chop_scale_tests,alias:chop_scale_tests"`
	ID                            uuid.UUID            `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Date                          time.Time            `bun:"item_date" json:"date"`
	Patient                       *Patient             `bun:"rel:belongs-to" json:"patient"`
	PatientID                     uuid.NullUUID        `bun:"type:uuid" json:"patientId"`
	ChopScaleTestResults          ChopScaleTestResults `bun:"rel:has-many" json:"chopScaleTestResults"`
	ChopScaleTestResultsForDelete []uuid.UUID          `bun:"-" json:"chopScaleTestResultsForDelete"`
}

type ChopScaleTests []*ChopScaleTest

func (items ChopScaleTests) GetChopScaleTestResults() ChopScaleTestResults {
	itemsForGet := make(ChopScaleTestResults, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].ChopScaleTestResults...)
	}
	return itemsForGet
}

func (items ChopScaleTests) GetChopScaleTestResultsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].ChopScaleTestResultsForDelete...)
	}
	return itemsForGet
}

func (item *ChopScaleTest) SetIDForChildren() {
	for i := range item.ChopScaleTestResults {
		item.ChopScaleTestResults[i].ChopScaleTestID = item.ID
	}
}

func (items ChopScaleTests) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}
