package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type QuestionVariant struct {
	bun.BaseModel `bun:"question_variants,alias:question_variants"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	QuestionID    uuid.NullUUID `bun:"type:uuid" json:"questionId"`
	Question      *Question     `bun:"rel:belongs-to" json:"question"`
}

type QuestionVariants []*QuestionVariant

type QuestionVariantsWithCount struct {
	QuestionVariants QuestionVariants `json:"items"`
	Count            int              `json:"count"`
}
