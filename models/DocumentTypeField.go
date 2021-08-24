package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentTypeField struct {
	bun.BaseModel  `bun:"document_type_fields,alias:document_type_fields"`
	ID             uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name           string    `json:"name"`
	Order          uint      `json:"order"`
	Type           string    `json:"type"`
	DocumentTypeID uuid.UUID `bun:"type:uuid" json:"documentTypeID"`
}
