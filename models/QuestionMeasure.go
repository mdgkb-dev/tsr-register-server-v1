package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type QuestionMeasure struct {
	bun.BaseModel `bun:"question_measures,alias:question_measures"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
	Order         int       `bun:"item_order" json:"order"`
	Question      *Question `bun:"rel:belongs-to" json:"question"`
	QuestionID    uuid.UUID `bun:"type:uuid" json:"questionId"`
}

type QuestionMeasures []*QuestionMeasure
