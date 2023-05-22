package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugDecrease struct {
	bun.BaseModel `bun:"drug_decreases,alias:drug_decreases"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	DrugArrive    *DrugArrive   `bun:"rel:belongs-to" json:"drugArrive"`
	DrugArriveID  uuid.NullUUID `bun:"type:uuid" json:"drugArriveId"`

	Quantity int        `json:"quantity"`
	Date     *time.Time `bun:"item_date" json:"date"`

	Patient   *Patient      `bun:"rel:belongs-to" json:"patient"`
	PatientID uuid.NullUUID `bun:"type:uuid" json:"patientId"`
	Comment   string        `json:"comment"`
}

type DrugDecreases []*DrugDecrease
type DrugDecreasesWithCount struct {
	DrugDecreases DrugDecreases `json:"items"`
	Count         int           `json:"count"`
}
