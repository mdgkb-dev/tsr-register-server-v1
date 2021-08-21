package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Document struct {
	bun.BaseModel  `bun:"document,alias:document"`
	ID             uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	DocumentType   *DocumentType `bun:"rel:belongs-to" json:"documentType"`
	DocumentTypeID uuid.UUID     `bun:"type:uuid" json:"documentTypeId"`
	Human          *Human        `bun:"rel:has-one" json:"human"`
	HumanID        uuid.UUID     `bun:"type:uuid" json:"humanId"`

	DocumentFieldValues []*DocumentFieldValue `bun:"rel:has-many" json:"documentFieldValues"`
}

func (item *Document) SetIdForChildren() {
	if len(item.DocumentFieldValues) > 0 {
		for i := range item.DocumentFieldValues {
			item.DocumentFieldValues[i].DocumentID = item.ID
		}
	}
}

func GetDocumentsFiledValues(docs []*Document) []*DocumentFieldValue {
	items := make([]*DocumentFieldValue, 0)
	if len(docs) == 0 {
		return items
	}
	for i := range docs {
		docs[i].SetIdForChildren()
		items = append(items, docs[i].DocumentFieldValues...)
	}
	return items
}
