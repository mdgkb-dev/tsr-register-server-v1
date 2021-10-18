package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Contact struct {
	bun.BaseModel `bun:"contact,alias:contact"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Phone         string        `json:"phone"`
	Email         string        `json:"email"`
	DeletedAt     time.Time     `bun:",soft_delete" json:"deletedAt"`
}
