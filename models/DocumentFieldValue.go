package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentFieldValue struct {
	bun.BaseModel `bun:"document_field_value,alias:document_field_value"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	ValueString   string    `json:"valueString"`
	ValueNumber   string    `json:"valueNumber"`
	ValueDate     string    `json:"valueDate"`

	Document   *Document `bun:"rel:has-one" json:"document"`
	DocumentID uuid.UUID `bun:"type:uuid" json:"documentId"`

	DocumentTypeField   *DocumentTypeField `bun:"rel:has-one" json:"documentTypeField"`
	DocumentTypeFieldID uuid.UUID          `bun:"type:uuid" json:"documentTypeFieldId"`
}
