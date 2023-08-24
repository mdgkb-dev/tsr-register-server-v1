package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type AnswerFile struct {
	bun.BaseModel `bun:"answer_files,alias:answer_files"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	FileInfoID    uuid.NullUUID `bun:"type:uuid" json:"fileInfoId"`
	FileInfo      *FileInfo     `bun:"rel:belongs-to" json:"fileInfo"`

	Comment string `json:"comment"`

	Answer   *Answer       `bun:"rel:belongs-to" json:"answer"`
	AnswerID uuid.NullUUID `bun:"type:uuid" json:"answerId"`
}

type AnswerFiles []*AnswerFile

type AnswerFilesWithCount struct {
	AnswerFiles AnswerFiles `json:"items"`
	Count       int         `json:"count"`
}

func (items AnswerFiles) SetFilePath(fileID *string) *string {
	for i := range items {
		path := items[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}

func (item *AnswerFile) SetFilePath(fileID *string) *string {
	if item.FileInfo.ID.UUID.String() == *fileID {
		item.FileInfo.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.FileInfo.FileSystemPath
	}
	return nil
}

func (items AnswerFiles) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)
	if len(items) == 0 {
		return itemsForGet
	}
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].FileInfo)
	}

	return itemsForGet
}

func (item *AnswerFile) SetForeignKeys() {
	item.FileInfoID = item.FileInfo.ID
}

func (items AnswerFiles) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}
