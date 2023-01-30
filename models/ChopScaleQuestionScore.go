package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ChopScaleQuestionScore struct {
	bun.BaseModel       `bun:"chop_scale_question_scores,alias:chop_scale_question_scores"`
	ID                  uuid.UUID          `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                string             `json:"name"`
	Score               int                `json:"score"`
	ChopScaleQuestion   *ChopScaleQuestion `bun:"rel:belongs-to" json:"chopScaleQuestion"`
	ChopScaleQuestionID uuid.UUID          `bun:"type:uuid" json:"chopScaleQuestionId"`
}

type ChopScaleQuestionScores []*ChopScaleQuestionScore
