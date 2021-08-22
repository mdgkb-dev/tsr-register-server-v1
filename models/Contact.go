package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Contact struct {
	bun.BaseModel `bun:"contact,alias:contact"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
}
