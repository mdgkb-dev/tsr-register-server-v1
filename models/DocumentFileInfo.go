package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentFileInfo struct {
	bun.BaseModel `bun:"document_file_infos,alias:document_file_infos"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	FileInfo      *FileInfo     `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoID    uuid.NullUUID `bun:"type:uuid" json:"fileInfoId"`
	Document      *Document     `bun:"rel:belongs-to" json:"document"`
	DocumentID    uuid.NullUUID `bun:"type:uuid" json:"documentId"`
}

type DocumentFileInfos []*DocumentFileInfo

type DocumentFileInfosWithCount struct {
	DocumentFileInfos DocumentFileInfos `json:"items"`
	Count             int               `json:"count"`
}

func GetFileInfoFileInfoToDocument(items []*DocumentFileInfo) []*FileInfo {
	itemsForGet := make([]*FileInfo, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].FileInfo)
	}
	return itemsForGet
}

func (item *Document) SetIDForChildren() {
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
		docs[i].SetIDForChildren()
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
		for j := range items[i].DocumentFileInfos {
			itemsForGet = append(itemsForGet, items[i].DocumentFileInfos[j].FileInfo)
		}
	}
	return itemsForGet
}
