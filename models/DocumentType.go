package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentType struct {
	bun.BaseModel      `bun:"document_types,alias:document_types"`
	ID                 uuid.NullUUID        `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name               string               `json:"name"`
	DocumentTypeFields []*DocumentTypeField `bun:"rel:has-many" json:"documentTypeFields"`
}

type DocumentTypes []*DocumentType

type DocumentTypesWithCount struct {
	DocumentTypes DocumentTypes `json:"items"`
	Count         int           `json:"count"`
}

func (item *DocumentType) SetIDForChildren() {
	if len(item.DocumentTypeFields) == 0 {
		return
	}
	for i := range item.DocumentTypeFields {
		item.DocumentTypeFields[i].DocumentTypeID = item.ID
	}
}
