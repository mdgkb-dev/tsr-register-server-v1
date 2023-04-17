package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type QuestionVariant struct {
	bun.BaseModel `bun:"questions_variants,alias:questions_variants"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	QuestionID    uuid.UUID     `bun:"type:uuid" json:"questionId"`
}

type QuestionVariants []*QuestionVariant
