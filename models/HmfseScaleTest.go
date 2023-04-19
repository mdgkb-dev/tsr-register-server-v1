package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type HmfseScaleTest struct {
	bun.BaseModel                  `bun:"hmfse_scale_tests,alias:hmfse_scale_tests"`
	ID                             uuid.UUID             `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Date                           time.Time             `bun:"item_date" json:"date"`
	Patient                        *Patient              `bun:"rel:belongs-to" json:"patient"`
	PatientID                      uuid.NullUUID         `bun:"type:uuid" json:"patientId"`
	HmfseScaleTestResults          HmfseScaleTestResults `bun:"rel:has-many" json:"hmfseScaleTestResults"`
	HmfseScaleTestResultsForDelete []uuid.UUID           `bun:"-" json:"hmfseScaleTestResultsForDelete"`
}

type HmfseScaleTests []*HmfseScaleTest

func (items HmfseScaleTests) GetHmfseScaleTestResults() HmfseScaleTestResults {
	itemsForGet := make(HmfseScaleTestResults, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].HmfseScaleTestResults...)
	}
	return itemsForGet
}

func (items HmfseScaleTests) GetHmfseScaleTestResultsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].HmfseScaleTestResultsForDelete...)
	}
	return itemsForGet
}

func (item *HmfseScaleTest) SetIDForChildren() {
	for i := range item.HmfseScaleTestResults {
		item.HmfseScaleTestResults[i].HmfseScaleTestID = item.ID
	}
}

func (items HmfseScaleTests) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}
