package models

import (
	"github.com/google/uuid"
	basemodels "github.com/pro-assistance/pro-assister/models"
	"github.com/uptrace/bun"
)

type DocumentTypeField struct {
	bun.BaseModel `bun:"document_type_fields,alias:document_type_fields"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Code          string        `json:"code"`
	Order         uint          `bun:"item_order"  json:"order"`

	ValueType   *basemodels.ValueType `bun:"rel:belongs-to" json:"valueType"`
	ValueTypeID uuid.NullUUID         `bun:"type:uuid" json:"valueTypeId"`

	DocumentType   *DocumentType `bun:"rel:belongs-to" json:"documentType"`
	DocumentTypeID uuid.NullUUID `bun:"type:uuid" json:"documentTypeID"`
}

type DocumentTypeFields []*DocumentTypeField
