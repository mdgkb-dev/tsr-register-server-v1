package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Anthropometry struct {
	bun.BaseModel `bun:"anthropometry,alias:anthropometry"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
	Measure       string    `json:"measure"`
}
