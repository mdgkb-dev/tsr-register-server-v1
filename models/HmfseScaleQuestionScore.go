package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type HmfseScaleQuestionScore struct {
	bun.BaseModel        `bun:"hmfse_scale_question_scores,alias:hmfse_scale_question_scores"`
	ID                   uuid.UUID           `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                 string              `json:"name"`
	Score                int                 `json:"score"`
	HmfseScaleQuestion   *HmfseScaleQuestion `bun:"rel:belongs-to" json:"hmfseScaleQuestion"`
	HmfseScaleQuestionID uuid.UUID           `bun:"type:uuid" json:"hmfseScaleQuestionId"`
}

type HmfseScaleQuestionScores []*HmfseScaleQuestionScore
