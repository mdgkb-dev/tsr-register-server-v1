package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FileInfoToDocument struct {
	bun.BaseModel `bun:"file_info_to_document,alias:file_info_to_document"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	FileInfo      *FileInfo `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoID    uuid.UUID `bun:"type:uuid" json:"fileInfoId"`
	Document      *Document `bun:"rel:belongs-to" json:"document"`
	DocumentID    uuid.UUID `bun:"type:uuid" json:"documentId"`
}

func GetFileInfoFileInfoToDocument(items []*FileInfoToDocument) []*FileInfo {
	itemsForGet := make([]*FileInfo, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].FileInfo)
	}
	return itemsForGet
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

func GetFileInfosFromDocuments(items []*Document) []*FileInfo {
	itemsForGet := make([]*FileInfo, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		for j := range items[i].FileInfoToDocument {
			itemsForGet = append(itemsForGet, items[i].FileInfoToDocument[j].FileInfo)
		}
	}
	return itemsForGet
}
