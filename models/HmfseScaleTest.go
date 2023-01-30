package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type HmfseScaleTest struct {
	bun.BaseModel `bun:"hmfse_scale_tests,alias:hmfse_scale_tests"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Date          time.Time `bun:"item_date" json:"date"`
	Patient       *Patient  `bun:"rel:belongs-to" json:"patient"`
	PatientID     uuid.UUID `bun:"type:uuid" json:"patientId"`
}

type HmfseScaleTests []*HmfseScaleTest
