package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RepresentativeType struct {
	bun.BaseModel  `bun:"representative_types,alias:representative_types"`
	ID             uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name           string    `json:"name"`
	ChildMaleType  string    `json:"childMaleType"`
	ChildWomanType string    `json:"childWomanType"`
	IsMale         bool      `json:"isMale"`
}
