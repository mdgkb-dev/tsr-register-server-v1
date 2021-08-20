package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Edv struct {
	bun.BaseModel `bun:"edv,alias:edv"`
	ID            uuid.UUID   `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string      `json:"name"`
	Disability    *Disability `bun:"rel:belongs-to" json:"disability"`
	DisabilityID  uuid.UUID   `bun:"type:uuid" json:"disabilityId"`
	Parameter1    bool        `json:"parameter1"`
	Parameter2    bool        `json:"parameter2"`
	Parameter3    bool        `json:"parameter3"`
}
