package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DrugApplicationFile struct {
	bun.BaseModel     `bun:"drug_application_files,alias:drug_application_files"`
	ID                uuid.NullUUID    `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name              string           `json:"name"`
	Comment           string           `json:"comment"`
	DrugApplication   *DrugApplication `bun:"rel:belongs-to" json:"drugApplication"`
	DrugApplicationID uuid.NullUUID    `bun:"type:uuid" json:"drugApplicationId"`
	FileInfo          *FileInfo        `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoID        uuid.NullUUID    `bun:"type:uuid" json:"fileInfoId"`
}

type DrugApplicationFiles []*DrugApplicationFile
type DrugApplicationFilesWithCount struct {
	DrugApplicationFiles DrugApplicationFiles `json:"items"`
	Count                int                  `json:"count"`
}
