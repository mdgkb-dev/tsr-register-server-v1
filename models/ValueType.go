package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ValueType struct {
	bun.BaseModel `bun:"value_type,alias:value_type"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
	ValueRelation string    `json:"valueRelation"`
}
