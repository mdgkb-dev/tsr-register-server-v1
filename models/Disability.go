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
}
