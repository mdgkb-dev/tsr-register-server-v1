package models

import (
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Document struct {
	bun.BaseModel  `bun:"document,alias:document"`
	ID             uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	DocumentType   *DocumentType `bun:"rel:belongs-to" json:"documentType"`
	DocumentTypeID uuid.UUID     `bun:"type:uuid" json:"documentTypeId"`
	Human          *Human        `bun:"rel:has-one" json:"human"`
	HumanID        uuid.UUID     `bun:"type:uuid" json:"humanId"`
	DeletedAt      *time.Time    `bun:",soft_delete" json:"deletedAt"`

	DocumentFieldValues          []*DocumentFieldValue `bun:"rel:has-many" json:"documentFieldValues"`
	DocumentFieldValuesForDelete []uuid.UUID           `bun:"-" json:"documentFieldValuesForDelete"`
	FileInfoToDocument           []*FileInfoToDocument `bun:"rel:has-many" json:"fileInfoToDocument"`
	FileInfoToDocumentForDelete  []uuid.UUID           `bun:"-" json:"fileInfoToDocumentForDelete"`
}

func (item *Document) SetFilePath(fileId *string) *string {
	for i := range item.FileInfoToDocument {
		if item.FileInfoToDocument[i].FileInfoID.String() == *fileId {
			item.FileInfoToDocument[i].FileInfo.FileSystemPath = uploadHelper.BuildPath(fileId)
			return &item.FileInfoToDocument[i].FileInfo.FileSystemPath
		}
	}
	return nil
}

func GetFileInfoToDocument(items []*Document) []*FileInfoToDocument {
	itemsForGet := make([]*FileInfoToDocument, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		items[i].SetIdForChildren()
		itemsForGet = append(itemsForGet, items[i].FileInfoToDocument...)
	}
	return itemsForGet
}

func GetFileInfoToDocumentForDelete(items []*Document) []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		items[i].SetIdForChildren()
		itemsForGet = append(itemsForGet, items[i].FileInfoToDocumentForDelete...)
	}
	return itemsForGet
}

func (item *Document) SetDeleteIdForChildren() {
	for i := range item.DocumentFieldValues {
		item.DocumentFieldValuesForDelete = append(item.DocumentFieldValuesForDelete, item.DocumentFieldValues[i].ID)
	}
	for i := range item.FileInfoToDocument {
		item.FileInfoToDocumentForDelete = append(item.FileInfoToDocumentForDelete, item.FileInfoToDocument[i].ID)
	}
}
