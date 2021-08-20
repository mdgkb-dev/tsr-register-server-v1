package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentType struct {
	bun.BaseModel               `bun:"document_types,alias:document_types"`
	ID                          uuid.UUID            `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                        string               `json:"name"`
	DocumentTypeFields          []*DocumentTypeField `bun:"rel:has-many" json:"documentTypeFields"`
	DocumentTypeFieldsForDelete []string             `bun:"-" json:"documentTypeFieldsForDelete"`
}

func (item *DocumentType) SetIdForChildren() {
	if len(item.DocumentTypeFields) == 0 {
		return
	}
	for i := range item.DocumentTypeFields {
		item.DocumentTypeFields[i].DocumentTypeID = item.ID
	}
}
