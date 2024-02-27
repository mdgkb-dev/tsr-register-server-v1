package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helpers/uploader"
	"github.com/uptrace/bun"
)

type Document struct {
	bun.BaseModel  `bun:"documents,alias:documents"`
	ID             uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	DocumentType   *DocumentType `bun:"rel:belongs-to" json:"documentType"`
	DocumentTypeID uuid.NullUUID `bun:"type:uuid" json:"documentTypeId"`
	Human          *Human        `bun:"rel:has-one" json:"human"`
	HumanID        uuid.NullUUID `bun:"type:uuid" json:"humanId"`
	DeletedAt      *time.Time    `bun:",soft_delete" json:"deletedAt"`

	DocumentFieldValues DocumentFieldValues `bun:"rel:has-many" json:"documentFieldValues"`
	DocumentFileInfos   DocumentFileInfos   `bun:"rel:has-many" json:"documentFileInfos"`
}

type Documents []*Document

type DocumentsWithCount struct {
	Documents Documents `json:"items"`
	Count     int       `json:"count"`
}

func (item *Document) SetFilePath(fileID *string) *string {
	for i := range item.DocumentFileInfos {
		if item.DocumentFileInfos[i].FileInfoID.UUID.String() == *fileID {
			item.DocumentFileInfos[i].FileInfo.FileSystemPath = uploader.BuildPath(fileID)
			return &item.DocumentFileInfos[i].FileInfo.FileSystemPath
		}
	}
	return nil
}

func GetFileInfoToDocument(items []*Document) []*DocumentFileInfo {
	itemsForGet := make([]*DocumentFileInfo, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		items[i].SetIDForChildren()
		itemsForGet = append(itemsForGet, items[i].DocumentFileInfos...)
	}
	return itemsForGet
}
