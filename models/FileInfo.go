package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FileInfo struct {
	bun.BaseModel  `bun:"file_infos,alias:file_infos"`
	ID             uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	OriginalName   string    `json:"originalName"`
	FileSystemPath string    `json:"fileSystemPath"`
	DeletedAt      time.Time `bun:",soft_delete" json:"deletedAt"`
}
