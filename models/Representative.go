package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Representative struct {
	bun.BaseModel `bun:"representative,alias:representative"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `

	Human   *Human    `bun:"rel:belongs-to" json:"human"`
	HumanID uuid.UUID `bun:"type:uuid" json:"humanId"`
}
