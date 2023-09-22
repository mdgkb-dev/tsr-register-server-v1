package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type QuestionFilter struct {
	bun.BaseModel `bun:"questions_filters,alias:questions_filters"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	QuetionsID    uuid.NullUUID `bun:"type:uuid" json:"questionId"`
}

type QuestionsFilters []*QuestionFilter
