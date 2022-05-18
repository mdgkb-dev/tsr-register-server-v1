package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RepresentativeType struct {
	bun.BaseModel   `bun:"representative_types,alias:representative_types"`
	ID              uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	ParentMaleType  string    `json:"parentMaleType"`
	ParentWomanType string    `json:"parentWomanType"`
	ChildMaleType   string    `json:"childMaleType"`
	ChildWomanType  string    `json:"childWomanType"`
}
