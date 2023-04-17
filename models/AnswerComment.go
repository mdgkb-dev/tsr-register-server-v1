package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type AnswerComment struct {
	bun.BaseModel `bun:"answer_comments,alias:answer_comments"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`

	Question   *Question     `bun:"rel:belongs-to" json:"question"`
	QuestionID uuid.NullUUID `bun:"type:uuid" json:"questionId"`

	Answer   *AnswerVariant `bun:"rel:belongs-to" json:"answer"`
	AnswerID uuid.NullUUID  `bun:"type:uuid" json:"answerIs"`

	Order int `bun:"item_order" json:"order"`
}

type AnswerComments []*AnswerComment
