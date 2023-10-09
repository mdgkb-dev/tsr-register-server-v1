package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugDozeComponent struct {
	bun.BaseModel `bun:"drug_doze_components,alias:drug_doze_components"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Code          string        `json:"code"`

	Measure  string `json:"measure"`
	Quantity int    `json:"quantity"`

	DrugDoze   *DrugDoze     `bun:"rel:belongs-to" json:"drugDoze"`
	DrugDozeID uuid.NullUUID `bun:"type:uuid" json:"drugDozeId"`
}

type DrugDozeComponents []*DrugDozeComponent

type DrugComponentsWithCount struct {
	DrugDozeComponents DrugDozeComponents `json:"items"`
	Count              int                `json:"count"`
}
