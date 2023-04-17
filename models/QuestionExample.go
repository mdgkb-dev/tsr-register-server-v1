package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type QuestionExample struct {
	bun.BaseModel `bun:"question_examples,alias:question_examples"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
	QuestionID    uuid.UUID `bun:"type:uuid" json:"questionId"`
	Question      *Question `bun:"rel:belongs-to" json:"question"`
	Order         int       `bun:"item_order" json:"order"`
}

type QuestionExamples []*QuestionExample
