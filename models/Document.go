package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"path/filepath"
)

type Document struct {
	bun.BaseModel  `bun:"document,alias:document"`
	ID             uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	DocumentType   *DocumentType `bun:"rel:belongs-to" json:"documentType"`
	DocumentTypeID uuid.UUID     `bun:"type:uuid" json:"documentTypeId"`
	Human          *Human        `bun:"rel:has-one" json:"human"`
	HumanID        uuid.UUID     `bun:"type:uuid" json:"humanId"`

	DocumentFieldValues []*DocumentFieldValue `bun:"rel:has-many" json:"documentFieldValues"`
	FileInfoToDocument  []*FileInfoToDocument `bun:"rel:has-many" json:"fileInfoToDocument"`
}

func (item *Document) SetFilePath(fileCategory *string, path *string) {
	for i := range item.FileInfoToDocument {
		fileId := item.FileInfoToDocument[i].FileInfo.ID.String()
		if fileId == *fileCategory {
			newPath := filepath.Join(*path, item.ID.String(), fileId)
			path = &newPath
			item.FileInfoToDocument[i].FileInfo.FileSystemPath = newPath
		}
	}
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

type FileInfoToDocument struct {
	bun.BaseModel `bun:"file_info_to_document,alias:file_info_to_document"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
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
