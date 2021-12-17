package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RegisterUser struct {
	bun.BaseModel      `bun:"registers_users,alias:registers_user"`
	ID                 uuid.UUID         `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Register   *Register `bun:"rel:belongs-to" json:"register"`
	RegisterID uuid.UUID         `bun:"type:uuid" json:"registerId"`
	User               *User             `bun:"rel:belongs-to" json:"user"`
	UserID             uuid.UUID         `bun:"type:uuid" json:"userId"`
}

type RegistersUsers []*RegisterUser
